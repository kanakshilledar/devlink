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

var usersCollection *mongo.Collection // users DATABASE
var secretKey = []byte(os.Getenv("KEY"))
var signedToken string

const USERS = "Users"

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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
