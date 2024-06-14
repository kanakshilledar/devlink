package controller

import (
	"context"
	"devlink/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

var collection *mongo.Collection

var collection2 *mongo.Collection

const database = "devlink"

const colName = "Users"

const colName2 = "Events"

func ConnDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientoption := options.Client().ApplyURI(os.Getenv("CONN"))
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully to the Cluster !!")

	collection = client.Database(database).Collection(colName)

	collection2 = client.Database(database).Collection(colName2)

	return client
}

func insertUser(user models.User) *mongo.InsertOneResult {
	user.Id = primitive.NewObjectID()
	insertone, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", insertone.InsertedID)
	return insertone
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	response := insertUser(user)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func insertEvent(event models.EventInfo) *mongo.InsertOneResult {
	event.EventId = primitive.NewObjectID()
	insertone, err := collection2.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted a single document: ", insertone.InsertedID)
	return insertone
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var event models.EventInfo

	_ = json.NewDecoder(r.Body).Decode(&event)
	response := insertEvent(event)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
