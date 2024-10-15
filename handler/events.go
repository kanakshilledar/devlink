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
	"log"
	"os"
)

var eventsCollection *mongo.Collection // events DATABASE

const DATABASE = "devlink"
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
	eventsCollection = client.Database(DATABASE).Collection(EVENTS)
}

func fetchUserIDFromEmail(email string) (primitive.ObjectID, error) {
	var user models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, errors.New("user not found")
	}
	return user.Id, nil
}

func FetchUserNameFromEmail(email string) (string, error) {
	var user models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}
	return user.Name, nil
}

func CheckEventExists(eventName string) (bool, error) {
	filter := bson.D{
		{Key: "eventName", Value: eventName},
	}
	var existingEvent models.EventInfo
	err := eventsCollection.FindOne(context.TODO(), filter).Decode(&existingEvent)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func InsertEvent(event models.EventInfo, userEmail string) (interface{}, error) {
	// fetch userid from email
	userID, err := fetchUserIDFromEmail(userEmail)
	if err != nil {
		return nil, err
	}

	userName, err := FetchUserNameFromEmail(userEmail)
	if err != nil {
		return nil, err
	}

	event.EventId = primitive.NewObjectID()
	event.AddedByID = userID
	event.AddedByName = userName
	insertOne, err := eventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Inserted a single document: ", insertOne.InsertedID)
	return insertOne, nil
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

func GetUserEvents(userId string) []primitive.M {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"added_by": id}
	var results []primitive.M
	cursor, err := eventsCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var event bson.M
		err := cursor.Decode(&event)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, event)
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
