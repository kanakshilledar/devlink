package router

import (
	"devlink/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.LandingPage).Methods("GET")
	router.HandleFunc("/api/createUser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/createEvent", controller.CreateEvent).Methods("POST")
	return router
}
