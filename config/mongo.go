package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {

	log.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Ghandi_01:SuGu5EAmjQxRSND8@cluster0.xpobc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {

		log.Fatal("Error connecting to MongoDB:", err)

	}

	DB = client.Database("unairsatu_v2")
	log.Println("Connected to MongoDB!")

}
func GetCollection(collectionName string) *mongo.Collection {
	// ConnectDB()
	if DB == nil {
		// log.Fatal("Database connection is not initialized")
		ConnectDB()
	}
	return DB.Collection(collectionName)
}
