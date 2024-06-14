package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phone_number"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	Company     string             `json:"company"`
}

type EventInfo struct {
	EventId     primitive.ObjectID `json:"event_id" bson:"_id"`
	EventName   string             `json:"event_name"`
	StartDate   string             `json:"start_date"`
	EndDate     string             `json:"end_date"`
	Description string             `json:"description"`
	EventType   string             `json:"event_type"`
	Company     string             `json:"company"`
}
