/*
Package main implements the whole devlink backend.
It runs the backend on port `8080` and is based on the `net/http` package
and uses `gorilla/mux`.

Usage:
	Add the mongoDB connection URL and JWT token KEY in **.env** before running.
	`$ go run main.go`
*/

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
