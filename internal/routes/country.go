package routes

import (
	"diving-log-book-service/internal/controllers"
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func CountryRouteInit(router *mux.Router) {
	countryRepo := repositories.NewCountryRepository(db.DB)
	countryService := services.NewCountryService(countryRepo)
	controller := controllers.NewCountryController(countryService)
	group := router.PathPrefix("/country").Subrouter()

	group.HandleFunc("/list", controller.ReadAll).Methods(http.MethodGet)
}
