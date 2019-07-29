package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Quote struct {
	quote string
}

func main() {

	ctx := context.TODO()

	// The URL which can be used in Java Applications in used here.
	mongoURI := fmt.Sprintf("mongodb+srv://playground:pwd123@playgroundcluster-q6nq5.mongodb.net/test?retryWrites=true&w=majority")

	fmt.Println("connection string is:", mongoURI)

	// Set client options and connect
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("kanye_quotes").Collection("quotes")

	// Inserting Data in MongoDB

	insertResult, err := collection.InsertOne(ctx, bson.M{"quote": "I'm a pornstar."})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Inserted Data to DB", insertResult.InsertedID)

}
