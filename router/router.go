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
	router.HandleFunc("/api/updateEvent/{id}", controller.UpdateEventInfo).Methods("PUT")
	router.HandleFunc("/api/getAll", controller.GetEvents).Methods("GET")
	router.HandleFunc("/api/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", controller.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/secret", controller.Secret).Methods("GET")
	return router
}
