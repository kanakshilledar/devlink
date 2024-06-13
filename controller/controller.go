package controller

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func ConnDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	clientoption := options.Client().ApplyURI(os.Getenv("CONN"))

	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected Successfully to the Cluster !!")
	return client

}
