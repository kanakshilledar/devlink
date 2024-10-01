package router

import (
	"devlink/controller"
	"devlink/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/", controller.LandingPage).Methods("GET")
	router.HandleFunc("/api/createUser", controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", controller.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/event/all", controller.GetAllEventsHandler).Methods("GET")

	// Protected routes with JWT middleware
	router.Handle("/api/createEvent", middleware.JWTmiddleware(http.HandlerFunc(controller.CreateEventHandler))).Methods("POST")
	router.Handle("/api/updateEvent/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.UpdateEventHandler))).Methods("PUT")
	router.Handle("/api/deleteEvent/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.DeleteEventHandler))).Methods("DELETE")
	router.Handle("/api/user/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.GetUserHandler))).Methods("GET")
	router.Handle("/api/updateUser/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.UpdateUserHandler))).Methods("PUT")
	router.Handle("/api/event/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.GetOneEventHandler))).Methods("GET")
	router.Handle("/api/event/user/{id}", middleware.JWTmiddleware(http.HandlerFunc(controller.GetUserEventHandler))).Methods("GET")

	return router
}
