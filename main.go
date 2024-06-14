package main

import (
	"devlink/router"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var db *mongo.Client

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
