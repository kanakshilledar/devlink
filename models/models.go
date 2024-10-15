// Package models defines the data models used in the application.
// These models represent various entities like users, events, login data,
// and response structures, which are stored in a MongoDB database
// using BSON (Binary JSON) format.
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user of the application.
// Fields:
//   - Id: MongoDB ObjectID for the user, automatically generated.
//   - Name: The user's full name.
//   - PhoneNumber: The user's phone number.
//   - Email: The user's email address, used for login and contact.
//   - Password: The user's hashed password.
//   - Company: The name of the company the user is associated with.
type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Company     string             `json:"company" bson:"company"`
}

// EventInfo represents an event in the application.
// Fields:
//   - EventId: MongoDB ObjectID for the event, automatically generated.
//   - EventName: The name of the event (required).
//   - StartDate: The start date of the event.
//   - EndDate: The end date of the event.
//   - Description: A description of the event.
//   - EventType: The type/category of the event.
//   - EventLink: A link for the event (required).
//   - Company: The company hosting or related to the event.
//   - Location: The location where the event takes place.
//   - AddedByID: MongoDB ObjectID of the user who added the event.
//   - AddedByName: The name of the user who added the event.
type EventInfo struct {
	EventId     primitive.ObjectID `json:"eventID,omitempty" bson:"_id,omitempty"`
	EventName   string             `json:"eventName" bson:"eventName" validate:"required"`
	StartDate   string             `json:"startDate" bson:"startDate"`
	EndDate     string             `json:"endDate" bson:"endDate"`
	Description string             `json:"description" bson:"description"`
	EventType   string             `json:"eventType" bson:"eventType"`
	EventLink   string             `json:"eventLink" bson:"eventLink" validate:"required"`
	Company     string             `json:"company" bson:"company"`
	Location    string             `json:"location" bson:"location"`
	AddedByID   primitive.ObjectID `json:"addedByID" bson:"addedByID"`
	AddedByName string             `json:"addedByName" bson:"addedByName"`
}

// Login represents the data required for user login.
// Fields:
//   - Email: The email address of the user.
//   - Password: The user's password.
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response represents a standard API response structure.
// Fields:
//   - Success: Indicates whether the operation was successful or not.
//   - Message: A message detailing the result of the operation.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
