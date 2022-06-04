package models

import "go.mongodb.org/mongo-driver/mongo"

type App struct {
	MongoClient    *mongo.Client
	TodoCollection *mongo.Collection
}
