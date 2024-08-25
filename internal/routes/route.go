package routes

import (
	"diving-log-book-service/internal/middlewares"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	publicRouter := router.PathPrefix("/v1").Subrouter()
	InitPublicRoute(publicRouter)

	privateRouter := router.PathPrefix("/v1").Subrouter()
	privateRouter.Use(middlewares.LoggingMiddleware)
	InitPrivateRoute(privateRouter)
}

func InitPublicRoute(router *mux.Router) {
	HealthcheckoutRouteInit(router)
}

func InitPrivateRoute(router *mux.Router) {
	DiveRouteInit(router)
	StorageRouteInit(router)
	CountryRouteInit(router)
	IslandRouteInit(router)
}
