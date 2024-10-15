// Package router defines the application's routing configuration.
// It uses the Gorilla Mux package to define routes for both public
// and protected (JWT-secured) endpoints, allowing access to various
// API actions like user and event management.
package router

import (
	"devlink/controller"     // Controller package containing handler functions
	"devlink/middleware"     // Middleware package providing JWT authentication
	"github.com/gorilla/mux" // Gorilla Mux for routing
	"net/http"               // HTTP utilities
)

// Router initializes and returns a new Gorilla Mux router with defined routes.
// It defines both public and protected routes. Public routes are accessible
// without authentication, while protected routes require JWT-based authentication
// using middleware.
//
// Public Routes:
//   - GET "/"                      : Serves the landing page.
//   - POST "/api/createUser"        : Creates a new user account.
//   - POST "/api/login"             : Authenticates a user and issues a JWT.
//   - GET "/api/logout"             : Logs out the current user.
//   - GET "/api/event/all"          : Retrieves all events.
//
// Protected Routes (require JWT authentication):
//   - POST "/api/createEvent"       : Creates a new event.
//   - PUT "/api/updateEvent/{id}"   : Updates an existing event by ID.
//   - DELETE "/api/deleteEvent/{id}": Deletes an event by ID.
//   - GET "/api/user/{id}"          : Retrieves user information by ID.
//   - PUT "/api/updateUser/{id}"    : Updates user information by ID.
//   - GET "/api/event/{id}"         : Retrieves event details by ID.
//   - GET "/api/event/user/{id}"    : Retrieves events created by a specific user.
//
// Returns:
//   - *mux.Router: A router instance with the registered routes.
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
