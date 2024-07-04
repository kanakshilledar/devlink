package controller

import (
	"devlink/handler"
	"devlink/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
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

func GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	events := handler.GetAllEvents()
	err := json.NewEncoder(w).Encode(events)
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
