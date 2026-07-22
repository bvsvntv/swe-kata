package main

import (
	"log"
	"net/http"
	"os"

	"dvdrental-api/config"
	"dvdrental-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system env vars")
	}

	config.ConnectDB()
	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server listening at http://localhost:%v", port)

	srvErr := server.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
