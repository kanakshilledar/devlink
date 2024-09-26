package controller

import (
	"devlink/handler"
	"devlink/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var LoggedIn = false

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, token, err := handler.Login(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		log.Fatal(err)
		return
	}

	type LoginResponse struct {
		Success bool   `json:"success"`
		Token   string `json:"token,omitempty"`
	}

	if response {
		resp := LoginResponse{
			Success: true,
			Token:   token,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusUnauthorized)
			return
		}
	} else {
		resp := LoginResponse{
			Success: false,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusUnauthorized)
			return
		}
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	err := json.NewEncoder(w).Encode("[+] Logged Out Successfully")
	if err != nil {
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	user := handler.GetUser(params["id"])
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var info models.User
	_ = json.NewDecoder(r.Body).Decode(&info)
	err := handler.UpdateUser(params["id"], info)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode("[+] Updated data successfully")
	if err != nil {
		log.Println(err)
	}
}
