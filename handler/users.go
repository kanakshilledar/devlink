// Package handler manages user-related operations such as creating, logging in, and updating users.
// It interacts with the MongoDB database and handles JWT token generation for authentication.
package handler

import (
	"context"
	"devlink/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

// MongoDB collection to store users.
var usersCollection *mongo.Collection

// Secret key used for signing JWT tokens.
var secretKey = []byte(os.Getenv("KEY"))

// USERS Name of the MongoDB collection to store users.
const USERS = "Users"

// init initializes the handler by connecting to the MongoDB database and loading environment variables.
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientOption := options.Client().ApplyURI(os.Getenv("CONN"))
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] Connected Successfully to the Cluster !!")
	usersCollection = client.Database(DATABASE).Collection(USERS)

	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("KEY environment variable not set")
	}
	secretKey = []byte(key)
}

// hashPassword generates a hashed version of the given password using bcrypt.
//
// Parameters:
//   - password: the raw password to be hashed.
//
// Returns:
//   - string: the hashed password.
//   - error: an error if hashing fails.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// checkPasswordHash compares a raw password with its hashed version to verify if they match.
//
// Parameters:
//   - password: the raw password provided by the user.
//   - hash: the hashed password stored in the database.
//
// Returns:
//   - bool: true if the password matches the hash, false otherwise.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckUserExists checks if a user with the provided email already exists in the database.
//
// Parameters:
//   - email: the email to check for existence.
//
// Returns:
//   - bool: true if the user exists, false otherwise.
//   - error: any error encountered during the database query.
func CheckUserExists(email string) (bool, error) {
	filter := bson.D{
		{Key: "email", Value: email},
	}
	var results models.User
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// InsertUser inserts a new user into the MongoDB collection after hashing their password.
//
// Parameters:
//   - user: the user object containing user details to be inserted.
//
// Returns:
//   - *mongo.InsertOneResult: result of the insert operation, including the inserted document ID.
func InsertUser(user models.User) *mongo.InsertOneResult {
	user.Id = primitive.NewObjectID()
	//fmt.Printf("[+] Inserted User %T\n", user.Password)
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hashedPassword

	insertOne, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[+] Inserted a single document: %+v\n", insertOne.InsertedID)
	return insertOne
}

// createToken generates a signed JWT token for the user, using the provided email as the subject.
//
// Parameters:
//   - email: the email of the user for which to generate the token.
//
// Returns:
//   - string: the signed JWT token.
//   - error: any error encountered during the token generation.
func createToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "devlink",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	//fmt.Println("[+] Token claims added: %+v\n", claims)
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Login checks the provided login credentials, verifies the password, and generates a JWT token
// if the login is successful.
//
// Parameters:
//   - info: the login object containing the user's email and password.
//
// Returns:
//   - bool: true if the login is successful, false otherwise.
//   - string: the generated JWT token if the login is successful.
//   - error: any error encountered during the process.
func Login(info models.Login) (bool, string, error) {
	filter := bson.D{
		{Key: "email", Value: info.Email},
	}
	var results models.User

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("[-] No documents found")
			return false, "", err
		}
		fmt.Println(err)
		return false, "", err
	}
	//fmt.Println(results.Id.Hex())

	response := checkPasswordHash(info.Password, results.Password)
	if !response {
		fmt.Println("[!] Password mismatch")
		return false, "", nil
	}

	fmt.Println("[+] Login Successfully")
	signedToken, err := createToken(results.Email)
	if err != nil {
		return false, "", err
	}

	fmt.Printf("[*] Token Created: %s\n", signedToken)
	return response, signedToken, nil
}

// GetUser retrieves a user from the MongoDB collection based on their unique ID.
//
// Parameters:
//   - userId: the ID of the user to be retrieved.
//
// Returns:
//   - models.User: the user object corresponding to the given ID.
func GetUser(userId string) models.User {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var results models.User
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

// UpdateUser updates the details of a user in the MongoDB collection based on their unique ID.
//
// Parameters:
//   - userId: the ID of the user to be updated.
//   - user: the user object containing the updated information.
//
// Returns:
//   - error: any error encountered during the update process.
func UpdateUser(userId string, user models.User) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("[-] Cannot convert to ObjectID")
	}
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	update := bson.D{{"$set", bson.D{
		{"name", user.Name},
		{"phone_number", user.PhoneNumber},
		{"email", user.Email},
		{"company", user.Company},
	}}}
	res := usersCollection.FindOneAndUpdate(context.TODO(), filter, update)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return fmt.Errorf("[-] No Documents Found")
		}
		return res.Err()
	}
	fmt.Println("[+] Update Successfully")
	return nil
}
