package controllers

import (
	"context"
	"fmt"
	"log"
	"todo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient gets mongo client.
func GetClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client
}

// GetConnectionHandle returns a collection handle.
func GetConnectionHandle(client *mongo.Client, collection string) *mongo.Collection {
	mongoCollection := client.Database("todo-app").Collection(collection)

	return mongoCollection
}

// InsertDocument inserts a todo item to database.
func InsertDocument(collection *mongo.Collection, item models.Item) {
	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// DeleteDocument delets a todo item from database.
func DeleteDocument(collection *mongo.Collection, item primitive.ObjectID) {
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": item})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted a single document: ", deleteResult)
}

// GetAllDocuments gets all documents from the given collection.
func GetAllDocuments(collection *mongo.Collection) []models.Item {
	findOptions := options.Find()
	// findOptions.SetLimit(2)

	var results []models.Item

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Item
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}
