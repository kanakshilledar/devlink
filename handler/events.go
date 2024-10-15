// Package handler provides functionalities for managing events in a MongoDB database.
// This package includes methods for user authentication, event creation, retrieval, updating, and deletion.
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

// init initializes the MongoDB client and sets up the events collection.
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

// fetchUserIDFromEmail retrieves the user ID associated with the provided email.
//
// Parameters:
//   - email: the email of the user whose ID is to be fetched.
//
// Returns:
//   - primitive.ObjectID: the ID of the user.
//   - error: an error if the user is not found or if an issue occurs during the database query.
func fetchUserIDFromEmail(email string) (primitive.ObjectID, error) {
	var user models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, errors.New("user not found")
	}
	return user.Id, nil
}

// FetchUserNameFromEmail retrieves the user's name associated with the provided email.
//
// Parameters:
//   - email: the email of the user whose name is to be fetched.
//
// Returns:
//   - string: the name of the user.
//   - error: an error if the user is not found or if an issue occurs during the database query.
func FetchUserNameFromEmail(email string) (string, error) {
	var user models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}
	return user.Name, nil
}

// CheckEventExists checks if an event with the specified event name already exists.
//
// Parameters:
//   - eventName: the name of the event to check for existence.
//
// Returns:
//   - bool: true if the event exists, false otherwise.
//   - error: an error if the database query fails.
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

// InsertEvent adds a new event to the database and associates it with the user identified by the provided email.
//
// Parameters:
//   - event: the event to be inserted.
//   - userEmail: the email of the user adding the event.
//
// Returns:
//   - interface{}: the inserted document's ID.
//   - error: an error if the event cannot be inserted.
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

// GetAllEvents retrieves all events from the database.
//
// Returns:
//   - []primitive.M: a slice of all events.
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

// GetOneEvent retrieves a single event based on the provided event ID.
//
// Parameters:
//   - eventId: the ID of the event to retrieve.
//
// Returns:
//   - models.EventInfo: the requested event.
//   - error: an error if the event cannot be found.
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

// GetUserEvents retrieves all events added by a specific user based on user ID.
//
// Parameters:
//   - userId: the ID of the user whose events are to be retrieved.
//
// Returns:
//   - []primitive.M: a slice of events added by the user.
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

// UpdateEvent updates the details of an existing event based on the provided event ID.
//
// Parameters:
//   - eventID: the ID of the event to update.
//   - event: the updated event information.
//
// Returns:
//   - error: an error if the event cannot be updated.
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

// DeleteEvent removes an event from the database based on the provided event ID.
//
// Parameters:
//   - eventId: the ID of the event to delete.
//
// Returns:
//   - void: no return value, but logs an error if the event cannot be deleted.
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
