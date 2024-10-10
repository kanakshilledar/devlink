package main

import (
	"devlink/router"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var db *mongo.Client

func main() {
	r := router.Router()
	// Handle OPTIONS requests
	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
