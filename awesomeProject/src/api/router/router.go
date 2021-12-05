package router

import (
	"awesomeProject/src/api/router/routes"
	"github.com/gorilla/mux"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}

