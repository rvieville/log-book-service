package main

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/middlewares"
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connection()
	apihelper.InitValidator()
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apihelper.Error(w, apihelper.NotFound("Route not found"))
}

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.CHeadersnMiddleware)
	router.Use(middlewares.Cors)
	routes.InitRoutes(router)
	router.NotFoundHandler = http.HandlerFunc(Handle404)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), router)
}
