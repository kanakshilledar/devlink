package handler

import (
	"context"
	"devlink/models"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

var usersCollection *mongo.Collection  // users DATABASE
var eventsCollection *mongo.Collection // events DATABASE

const DATABASE = "devlink"
const USERS = "Users"
const EVENTS = "Events"

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

	usersCollection = client.Database(DATABASE).Collection(USERS)
	eventsCollection = client.Database(DATABASE).Collection(EVENTS)
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
	insertOne, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[+] Inserted a single document: %+v\n", insertOne.InsertedID)
	return insertOne
}

func InsertEvent(event models.EventInfo) *mongo.InsertOneResult {
	event.EventId = primitive.NewObjectID()
	insertOne, err := eventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Inserted a single document: ", insertOne.InsertedID)
	return insertOne
}

func Login(info models.Login) bool {
	filter := bson.D{
		{Key: "email", Value: info.Email},
	}
	var results models.User

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("[-] No documents found")
		}
		fmt.Println(err)
	}

	response := checkPasswordHash(info.Password, results.Password)
	if response {
		fmt.Println("[+] Login Successfully")
	}
	return response
}

func GetUser(userId string) models.User {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var results models.User
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func GetAllEvents() []primitive.M {
	cursor, err := eventsCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var events []primitive.M
	for cursor.Next(context.Background()) {
		var event bson.M
		err := cursor.Decode(&event)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())
	return events
}

func GetOneEvent(eventId string) models.EventInfo {
	id, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var results models.EventInfo
	err = eventsCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func UpdateEvent(eventID string, event models.EventInfo) error {

	id, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return fmt.Errorf("[-] Cannot convert to ObjectID")
	}

	filter := bson.D{
		{Key: "_id", Value: id},
	}

	update := bson.D{{"$set", bson.D{
		{Key: "event_name", Value: event.EventName},
		{Key: "start_date", Value: event.StartDate},
		{Key: "end_date", Value: event.EndDate},
		{Key: "description", Value: event.Description},
		{Key: "event_type", Value: event.EventType},
		{Key: "company", Value: event.Company},
	}}}

	res := eventsCollection.FindOneAndUpdate(context.TODO(), filter, update)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			fmt.Println("[-] No Documents Found")
		}
		fmt.Println(res.Err())
	}

	fmt.Println("[+] Update Successfully")

	return nil
}

func DeleteEvent(eventId string) {
	id, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	_, err = eventsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[+] Deleted event with id: %s\n", id)
}
