package routes

import (
	"diving-log-book-service/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthcheckoutRouteInit(router *mux.Router) {
	controller := controllers.NewHealthcheckController()

	router.HandleFunc("/healthcheck", controller.Alive).Methods(http.MethodGet)
}
