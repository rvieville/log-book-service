package routes

import (
	"diving-log-book-service/internal/controllers/healthcheck"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthcheckoutRouteInit(router *mux.Router) {
	controller := healthcheck.New()

	router.HandleFunc("/healthcheck", controller.Alive).Methods(http.MethodGet)
}
