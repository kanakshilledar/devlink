// Package controller handles the main API functionality, including creating users, login, logout,
// and user profile management. The controller interacts with the models and handler packages.
package controller

import (
	"devlink/handler"
	"devlink/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CreateUserHandler handles the creation of a new user. It checks if the user already exists
// and inserts the user into the database if they don't.
//
// This handler is used for the `/api/createUser` endpoint and expects a POST request with a JSON payload
// containing user information.
//
// Example JSON request body:
//
//	{
//	  "name": "John Doe",
//	  "email": "john.doe@example.com",
//	  "phoneNumber": "1234567890",
//	  "password": "password123",
//	  "company": "Tech Co."
//	}
//
// Response:
//   - Success: true if the user was created successfully.
//   - Message: "User created" or "User already exists" depending on the operation result.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	exist, err := handler.CheckUserExists(user.Email)
	var response models.Response
	if exist == true {
		response.Success = false
		response.Message = "User already exists"
	} else {
		_ = handler.InsertUser(user)
		response.Success = true
		response.Message = "User created"
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

// LoginHandler handles user login by validating credentials and generating a JWT token.
//
// This handler is used for the `/api/login` endpoint and expects a POST request with a JSON payload
// containing login credentials (email and password).
//
// Example JSON request body:
//
//	{
//	  "email": "john.doe@example.com",
//	  "password": "password123"
//	}
//
// Response:
//   - Success: true if login is successful.
//   - Token: A JWT token if the login is successful.
//   - Success: false if login fails.
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

// LogoutHandler handles user logout. It invalidates the current session and returns a success message.
//
// This handler is used for the `/api/logout` endpoint and expects a GET request.
//
// Response:
//   - Success: true.
//   - Message: "logout successful".
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	response := models.Response{
		Success: true,
		Message: "logout successfull",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// GetUserHandler retrieves user information based on the user ID provided in the URL.
//
// This handler is used for the `/api/user/{id}` endpoint and expects a GET request.
// The user ID is passed as a URL parameter.
//
// Response:
//   - Success: true if the user is found.
//   - User: The user's profile information (ID, name, email, phone number, company).
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	type Response struct {
		Success bool        `json:"success"`
		User    models.User `json:"user"`
	}

	params := mux.Vars(r)
	user := handler.GetUser(params["id"])
	response := Response{
		Success: true,
		User:    user,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateUserHandler updates user information based on the user ID provided in the URL.
//
// This handler is used for the `/api/updateUser/{id}` endpoint and expects a PUT request
// with a JSON payload containing the updated user information.
//
// Example JSON request body:
//
//	{
//	  "name": "Jane Doe",
//	  "email": "jane.doe@example.com",
//	  "phoneNumber": "9876543210",
//	  "company": "New Tech Co."
//	}
//
// Response:
//   - Success: true if the update was successful.
//   - Message: A message indicating the result of the update.
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var info models.User
	_ = json.NewDecoder(r.Body).Decode(&info)
	err := handler.UpdateUser(params["id"], info)

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
