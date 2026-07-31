package main

import (
	"log"
	"net/http"
	"time"

	"dvdrental-api/internal/actors"
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/categories"
	"dvdrental-api/internal/films"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

type application struct {
	config config
	db     *pgx.Conn
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, types.HeartbeatResponse{
			Heartbeat: time.Now().UnixMilli(),
		})
	})

	filmServcice := films.NewService(repo.New(app.db))
	filmHandler := films.NewHandler(filmServcice)
	r.Get("/films", filmHandler.FetchFilms)
	r.Get("/films/{filmID}", filmHandler.GetFilm)

	actorServcice := actors.NewService(repo.New(app.db))
	actorHandler := actors.NewHandler(actorServcice)
	r.Get("/actors", actorHandler.FetchActors)
	r.Get("/actors/{actorID}", actorHandler.GetActor)
	r.Post("/actors", actorHandler.CreateActor)
	r.Delete("/actors/{actorID}", actorHandler.DeleteActor)
	r.Put("/actors/{actorID}", actorHandler.UpdateActor)
	r.Patch("/actors/{actorID}", actorHandler.UpdateActorPartial)
	r.Get("/actors/{actorID}/films", actorHandler.FetchActorFilms)

	categoryServcice := categories.NewService(repo.New(app.db))
	categoryHandler := categories.NewHandler(categoryServcice)
	r.Get("/categories", categoryHandler.FetchCategories)
	r.Get("/categories/{categoryID}", categoryHandler.GetCategory)
	r.Post("/categories", categoryHandler.CreateCategory)
	r.Delete("/categories/{categoryID}", categoryHandler.DeleteCategory)
	r.Put("/categories/{categoryID}", categoryHandler.UpdateCategory)
	r.Patch("/categories/{categoryID}", categoryHandler.UpdateCategoryPartial)
	r.Get("/categories/{categoryID}/films", categoryHandler.FetchCategoryFilms)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         ":" + app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("listening at http://localhost:%v\n", app.config.addr)

	return srv.ListenAndServe()
}
