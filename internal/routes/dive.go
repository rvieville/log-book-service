package routes

import (
	"diving-log-book-service/internal/controllers/dive"
	"net/http"

	"github.com/gorilla/mux"
)

func DiveRouteInit(router *mux.Router) {
	controller := dive.New()

	group := router.PathPrefix("/dive").Subrouter()

	group.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
}
