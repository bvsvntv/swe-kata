package main

import (
	"log"
	"net/http"
	"time"

	"dvdrental-api/internal/actors"
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/films"

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

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good\n"))
	})

	filmServcice := films.NewService(repo.New(app.db))
	filmHandler := films.NewHandler(filmServcice)
	r.Get("/films", filmHandler.GetFilms)

	actorServcice := actors.NewService(repo.New(app.db))
	actorHandler := actors.NewHandler(actorServcice)
	r.Get("/actors", actorHandler.GetActors)

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
