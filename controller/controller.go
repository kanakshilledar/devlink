package controller

import (
	"devlink/handler"
	"devlink/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var collection *mongo.Collection
var collection2 *mongo.Collection

func LandingPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	response := handler.InsertUser(user)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var event models.EventInfo

	_ = json.NewDecoder(r.Body).Decode(&event)
	response := handler.InsertEvent(event)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
