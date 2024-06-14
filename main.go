package main

import (
	"context"
	"devlink/controller"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var db *mongo.Client

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/insert", controller.CreateUser)

	db := controller.ConnDB()

	err := db.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		panic(err)
	}
}
