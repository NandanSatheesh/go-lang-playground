package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type QuotesData []string

type Quote struct {
	quote string `json:"quote"`
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

	file, err := ioutil.ReadFile("quotes.json")

	if err != nil {
		fmt.Println(err)
	}

	var data QuotesData

	_ = json.Unmarshal([]byte(file), &data)

	// Inserting Data in MongoDB

	for i, quote := range data {
		fmt.Println(i, " -> ", quote)
		insertResult, err := collection.InsertOne(ctx, bson.M{"quote": quote})

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Inserted Data to DB", insertResult.InsertedID)
	}

}
