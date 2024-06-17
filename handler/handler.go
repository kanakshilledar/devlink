package handler

import (
	"context"
	"devlink/models"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var collection *mongo.Collection
var collection2 *mongo.Collection

const database = "devlink"
const usersCollection = "Users"
const eventsCollection = "Events"

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientOption := options.Client().ApplyURI(os.Getenv("CONN"))
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Connected Successfully to the Cluster !!")

	collection = client.Database(database).Collection(usersCollection)
	collection2 = client.Database(database).Collection(eventsCollection)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InsertUser(user models.User) *mongo.InsertOneResult {
	user.Id = primitive.NewObjectID()
	//fmt.Printf("[+] Inserted User %T\n", user.Password)
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hashedPassword
	insertOne, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", insertOne.InsertedID)
	return insertOne
}

func InsertEvent(event models.EventInfo) *mongo.InsertOneResult {
	event.EventId = primitive.NewObjectID()
	insertOne, err := collection2.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertOne.InsertedID)
	return insertOne
}

func Login(info models.Login) bool {
	filter := bson.D{
		{Key: "email", Value: info.Email},
	}
	var results models.User

	err := collection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		panic(err)
	}

	response := checkPasswordHash(info.Password, results.Password)
	if response {
		fmt.Println("Login Successfully")
	}
	return response
}
