package routes

import (
	fishmapping "diving-log-book-service/internal/controllers/fishMapping"
	"net/http"

	"github.com/gorilla/mux"
)

func FishMappingRouteInit(router *mux.Router) {
	controller := fishmapping.New()

	group := router.PathPrefix("/fish-mapping").Subrouter()

	group.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
}
