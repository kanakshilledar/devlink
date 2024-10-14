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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		MaxAge:           8000,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
