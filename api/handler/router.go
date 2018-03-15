package handler

import (
	"api/model"

	"github.com/gorilla/mux"
)

func NewRouter(routes model.Routes) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Func)
	}

	return router
}
