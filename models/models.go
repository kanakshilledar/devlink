package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Company     string             `json:"company" bson:"company"`
}

type EventInfo struct {
	EventId     primitive.ObjectID `json:"event_id,omitempty" bson:"_id,omitempty"`
	EventName   string             `json:"event_name" bson:"event_name"`
	StartDate   string             `json:"start_date" bson:"start_date"`
	EndDate     string             `json:"end_date" bson:"end_date"`
	Description string             `json:"description" bson:"description"`
	EventType   string             `json:"event_type" bson:"event_type"`
	Company     string             `json:"company" bson:"company"`
	AddedBy     primitive.ObjectID `json:"added_by" bson:"added_by"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
