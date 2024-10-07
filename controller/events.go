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
	response, err := handler.InsertEvent(event, userEmail)
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

//func Secret(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "auth-session")
//
//	// Check if user is authenticated
//	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
//		http.Error(w, "Forbidden", http.StatusForbidden)
//		return
//	}
//
//	// Print secret message
//	_, err := fmt.Fprintln(w, "flag{you_logged_in}")
//	if err != nil {
//		panic(err)
//	}
//}

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

	params := mux.Vars(r)
	event := handler.GetOneEvent(params["id"])
	err := json.NewEncoder(w).Encode(event)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	events := handler.GetUserEvents(params["id"])
	err := json.NewEncoder(w).Encode(events)
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
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode("[+] Updated data successfully")
	if err != nil {
		log.Println(err)
	}
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	handler.DeleteEvent(params["id"])
	err := json.NewEncoder(w).Encode("[+] Deleted Event Successfully")
	if err != nil {
		log.Fatal(err)
	}
}
