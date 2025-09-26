package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	Database *mongo.Database
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// checks db connection that is alive
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	Database = client.Database("library-db")
	fmt.Println("Connected to MongoDB")
}

func Disconnect(ctx context.Context) error {
	if err := client.Disconnect(ctx); err != nil {
		return err
	}
	fmt.Println("Disconnected from MongoDB")
	return nil
}
