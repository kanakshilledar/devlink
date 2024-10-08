package controller

import (
	"devlink/handler"
	"devlink/middleware"
	"devlink/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"strings"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		log.Fatal(err)
	}
}

func extractUserIDFromToken(r *http.Request) (string, error) {
	signingKey := []byte(os.Getenv("KEY"))
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("no token provided")
	} else if len(tokenString) > 6 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if email, ok := claims["sub"].(string); ok {
			fmt.Println("[+] email: ", email)
			return email, nil
		}
	}

	return "", errors.New("invalid token claims")
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	userEmail, err := extractUserIDFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	var event models.EventInfo
	_ = json.NewDecoder(r.Body).Decode(&event)
	exist, err := handler.CheckEventExists(event.EventName)
	var response models.Response
	if exist == true {
		response.Success = false
		response.Message = "Event already exists"
	} else {
		_, err := handler.InsertEvent(event, userEmail)
		if err != nil {
			response.Success = false
			response.Message = err.Error()
		}
		response.Success = true
		response.Message = "Event created"
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error inserting event:", err)
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	type Response struct {
		Success bool          `json:"success"`
		Name    string        `json:"name"`
		Events  []primitive.M `json:"events"`
	}

	events := handler.GetAllEvents()

	tokenString := r.Header.Get("Authorization")
	response := Response{}
	if tokenString != "" {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := middleware.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			response.Name = ""
		} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["sub"].(string)
			name, _ := handler.FetchUserNameFromEmail(email)
			response.Name = name
		} else {
			response.Name = ""
		}
	} else {
		response.Name = ""
	}

	response.Success = true
	response.Events = events

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetOneEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	type Response struct {
		Success bool             `json:"success"`
		Events  models.EventInfo `json:"events"`
	}

	params := mux.Vars(r)
	event := handler.GetOneEvent(params["id"])
	response := Response{
		Success: true,
		Events:  event,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	type Response struct {
		Success bool          `json:"success"`
		Events  []primitive.M `json:"events"`
	}

	params := mux.Vars(r)
	events := handler.GetUserEvents(params["id"])
	response := Response{
		Success: true,
		Events:  events,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	var event models.EventInfo
	_ = json.NewDecoder(r.Body).Decode(&event)
	params := mux.Vars(r)
	err := handler.UpdateEvent(params["id"], event)

	response := models.Response{}
	response.Success = true
	response.Message = "Updated data successfully"

	if err != nil {
		response.Success = false
		response.Message = err.Error()
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	handler.DeleteEvent(params["id"])

	response := models.Response{
		Success: true,
		Message: "Deleted Event successfully",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}
