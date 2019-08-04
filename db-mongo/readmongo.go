package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	QuoteData string `json:"quote"`
	Id        int    `json:"id"`
}

func GetClient() *mongo.Client {

	mongoURI := fmt.Sprintf("mongodb+srv://playground:pwd123@playgroundcluster-q6nq5.mongodb.net/test?retryWrites=true&w=majority")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {

	ctx := context.TODO()
	// Set client options and connect
	client := GetClient()

	err := client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")

	var Result Data

	collection := client.Database("quotes_production").Collection("quotes_db_prod")

	// filter := bson.D{{"id", 12}}
	err = collection.FindOne(ctx, bson.M{"id": int32(1)}).Decode(&Result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Result.QuoteData, Result.Id)
}
