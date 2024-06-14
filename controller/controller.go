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

const database = "devlink"

const colName = "Users"

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

	return client
}

func insertUser(user models.User) {
	user.Id = primitive.NewObjectID()
	insertone, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", insertone.InsertedID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	insertUser(user)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}
