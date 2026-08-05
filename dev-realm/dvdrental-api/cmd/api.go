package main

import (
	"log"
	"net/http"
	"time"

	"dvdrental-api/internal/actors"
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/addresses"
	"dvdrental-api/internal/categories"
	"dvdrental-api/internal/cities"
	"dvdrental-api/internal/countries"
	"dvdrental-api/internal/customers"
	"dvdrental-api/internal/films"
	"dvdrental-api/internal/inventory"
	"dvdrental-api/internal/languages"
	"dvdrental-api/internal/rentals"
	"dvdrental-api/internal/staffs"
	"dvdrental-api/internal/stores"
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
	r.Get("/films/{filmID}/actors", filmHandler.FetchFilmActors)
	r.Get("/films/{filmID}/categories", filmHandler.FetchFilmCategories)
	r.Get("/films/{filmID}/inventory", filmHandler.FetchFilmInventory)

	actorServcice := actors.NewService(repo.New(app.db))
	actorHandler := actors.NewHandler(actorServcice)
	r.Get("/actors", actorHandler.FetchActors)
	r.Get("/actors/{actorID}", actorHandler.GetActor)
	r.Post("/actors", actorHandler.CreateActor)
	r.Delete("/actors/{actorID}", actorHandler.DeleteActor)
	r.Put("/actors/{actorID}", actorHandler.UpdateActor)
	r.Patch("/actors/{actorID}", actorHandler.UpdateActorPartial)
	r.Get("/actors/{actorID}/films", actorHandler.FetchActorFilms)

	categoryService := categories.NewService(repo.New(app.db))
	categoryHandler := categories.NewHandler(categoryService)
	r.Get("/categories", categoryHandler.FetchCategories)
	r.Get("/categories/{categoryID}", categoryHandler.GetCategory)
	r.Post("/categories", categoryHandler.CreateCategory)
	r.Delete("/categories/{categoryID}", categoryHandler.DeleteCategory)
	r.Put("/categories/{categoryID}", categoryHandler.UpdateCategory)
	r.Patch("/categories/{categoryID}", categoryHandler.UpdateCategoryPartial)
	r.Get("/categories/{categoryID}/films", categoryHandler.FetchCategoryFilms)

	languageService := languages.NewService(repo.New(app.db))
	languageHandler := languages.NewHandler(languageService)
	r.Get("/languages", languageHandler.FetchLanguages)
	r.Get("/languages/{languageID}", languageHandler.GetLanguage)

	countryService := countries.NewService(repo.New(app.db))
	countryHandler := countries.NewHandler(countryService)
	r.Get("/countries", countryHandler.FetchCountries)
	r.Get("/countries/{countryID}", countryHandler.GetCountry)
	r.Post("/countries", countryHandler.CreateCountry)
	r.Put("/countries/{countryID}", countryHandler.UpdateCountry)
	r.Delete("/countries/{countryID}", countryHandler.DeleteCountry)
	r.Get("/countries/{countryID}/cities", countryHandler.FetchCountryCities)

	cityService := cities.NewService(repo.New(app.db))
	cityHandler := cities.NewHandler(cityService)
	r.Get("/cities", cityHandler.FetchCities)
	r.Get("/cities/{cityID}", cityHandler.GetCity)
	r.Post("/cities", cityHandler.CreateCity)
	r.Delete("/cities/{cityID}", cityHandler.DeleteCity)
	r.Put("/cities/{cityID}", cityHandler.UpdateCity)
	r.Patch("/cities/{cityID}", cityHandler.UpdateCityPartial)

	addressService := addresses.NewService(repo.New(app.db))
	addressHandler := addresses.NewHandler(addressService)
	r.Get("/addresses", addressHandler.FetchAddresses)
	r.Get("/addresses/{addressID}", addressHandler.GetAddress)
	r.Post("/addresses", addressHandler.CreateAddress)
	r.Put("/addresses/{addressID}", addressHandler.UpdateAddress)
	r.Patch("/addresses/{addressID}", addressHandler.UpdateAddressPartial)
	r.Delete("/addresses/{addressID}", addressHandler.DeleteAddress)

	storeService := stores.NewService(repo.New(app.db))
	storeHandler := stores.NewHandler(storeService)
	r.Get("/stores", storeHandler.FetchStores)
	r.Get("/stores/{storeID}", storeHandler.GetStore)
	r.Post("/stores", storeHandler.CreateStore)
	r.Put("/stores/{storeID}", storeHandler.UpdateStore)
	r.Patch("/stores/{storeID}", storeHandler.UpdateStorePartial)
	r.Delete("/stores/{storeID}", storeHandler.DeleteStore)
	r.Get("/stores/{storeID}/inventory", storeHandler.FetchStoreInventory)
	r.Get("/stores/{storeID}/customers", storeHandler.FetchStoreCustomers)

	staffService := staffs.NewService(repo.New(app.db))
	staffHandler := staffs.NewHandler(staffService)
	r.Get("/staffs", staffHandler.FetchStaffs)
	r.Get("/staffs/{staffID}", staffHandler.GetStaff)
	r.Post("/staffs", staffHandler.CreateStaff)
	r.Put("/staffs/{staffID}", staffHandler.UpdateStaff)
	r.Patch("/staffs/{staffID}", staffHandler.UpdateStaffPartial)
	r.Delete("/staffs/{staffID}", staffHandler.DeleteStaff)

	inventoryService := inventory.NewService(repo.New(app.db))
	inventoryHandler := inventory.NewHandler(inventoryService)
	r.Get("/inventory", inventoryHandler.FetchInventories)
	r.Get("/inventory/{inventoryID}", inventoryHandler.GetInventory)
	r.Post("/inventory", inventoryHandler.CreateInventory)
	r.Delete("/inventory/{inventoryID}", inventoryHandler.DeleteInventory)

	customerService := customers.NewService(repo.New(app.db))
	customerHandler := customers.NewHandler(customerService)
	r.Get("/customers", customerHandler.FetchCustomers)
	r.Get("/customers/{customerID}", customerHandler.GetCustomer)
	r.Post("/customers", customerHandler.CreateCustomer)
	r.Put("/customers/{customerID}", customerHandler.UpdateCustomer)
	r.Patch("/customers/{customerID}", customerHandler.UpdateCustomerPartial)
	r.Delete("/customers/{customerID}", customerHandler.DeleteCustomer)

	rentalService := rentals.NewService(repo.New(app.db))
	rentalHandler := rentals.NewHandler(rentalService)
	r.Get("/rentals", rentalHandler.FetchRentals)
	r.Get("/rentals/{rentalID}", rentalHandler.GetRental)
	r.Post("/rentals", rentalHandler.CreateRental)
	r.Put("/rentals/{rentalID}", rentalHandler.UpdateRental)
	r.Delete("/rentals/{rentalID}", rentalHandler.DeleteRental)
	r.Post("/rentals/{rentalID}/return", rentalHandler.ReturnRental)

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
