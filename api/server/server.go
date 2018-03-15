package server

import (
	"api/controller"
	"api/handler"
	"api/model"
	"api/store"
	"log"
	"net/http"
)

func ListenAndServeTLS(addr string) {
	key := "./server/ssl/server.key"
	crt := "./server/ssl/server.cert"
	if err := http.ListenAndServeTLS(addr, crt, key, getRouter()); err != nil {
		log.Fatal(err)
	}
}

var routes model.Routes

func getRouter() http.Handler {
	defaultController := controller.NewDefault()
	routes = append(routes, defaultController.Routes...)

	// Replace store.MockType by store.MongoType
	manager, err := store.CreateManager(store.MongoType)
	if err != nil {
		log.Fatal(err)
	}

	taskControler := controller.NewTask(manager)
	routes = append(routes, taskControler.Routes...)

	return handler.NewRouter(routes)
}
