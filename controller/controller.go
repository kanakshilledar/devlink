package controller

import (
	"devlink/handler"
	"devlink/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var LoggedIn = false
var collection *mongo.Collection
var collection2 *mongo.Collection

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	response := handler.InsertUser(user)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var event models.EventInfo

	_ = json.NewDecoder(r.Body).Decode(&event)
	response := handler.InsertEvent(event)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	session, _ := store.Get(r, "auth-session")

	var user models.Login
	_ = json.NewDecoder(r.Body).Decode(&user)

	response := handler.Login(user)
	success := "[+] Login Success!"
	failure := "[+] Login Failure!"
	if response {
		session.Values["authenticated"] = true
		err := session.Save(r, w)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(success)
		if err != nil {
			panic(err)
		}
	}
	if !response {
		err := json.NewEncoder(w).Encode(failure)
		if err != nil {
			panic(err)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	session, _ := store.Get(r, "auth-session")
	//session.Options.MaxAge =
	session.Values["authenticated"] = false
	err := session.Save(r, w)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode("[+] Logged Out Successfully")
	if err != nil {
		panic(err)
	}
}

func Secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	_, err := fmt.Fprintln(w, "flag{you_logged_in}")
	if err != nil {
		panic(err)
	}
}
