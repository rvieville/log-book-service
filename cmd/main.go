package main

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/middlewares"
	"diving-log-book-service/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connection()
}

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.CHeadersnMiddleware)
	routes.InitRoutes(router)

	http.ListenAndServe(":8080", router)
}
