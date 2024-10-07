package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Company     string             `json:"company" bson:"company"`
}

type EventInfo struct {
	EventId     primitive.ObjectID `json:"eventID,omitempty" bson:"_id,omitempty"`
	EventName   string             `json:"eventName" bson:"eventName" validate:"required"`
	StartDate   string             `json:"startDate" bson:"startDate"`
	EndDate     string             `json:"endDate" bson:"endDate"`
	Description string             `json:"description" bson:"description"`
	EventType   string             `json:"eventType" bson:"eventType"`
	EventLink   string             `json:"eventLink" bson:"eventLink" validate:"required"`
	Company     string             `json:"company" bson:"company"`
	AddedByID   primitive.ObjectID `json:"addedByID" bson:"addedByID"`
	AddedByName string             `json:"addedByName" bson:"addedByName"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
