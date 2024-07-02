package router

import (
	"devlink/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.LandingPage).Methods("GET")

	router.HandleFunc("/api/createUser", controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/createEvent", controller.CreateEventHandler).Methods("POST")
	router.HandleFunc("/api/updateEvent/{id}", controller.UpdateEventHandler).Methods("PUT")
	router.HandleFunc("/api/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", controller.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/secret", controller.Secret).Methods("GET")
	router.HandleFunc("/api/deleteEvent/{id}", controller.DeleteEventHandler).Methods("DELETE")
	router.HandleFunc("/api/user/{id}", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/updateUser/{id}", controller.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/event/all", controller.GetAllEventsHandler).Methods("GET")
	router.HandleFunc("/api/event/{id}", controller.GetOneEventHandler).Methods("GET")

	return router
}
