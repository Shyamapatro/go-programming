// config/database.go
package config

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB connects to MongoDB and returns a client instance
func ConnectMongoDB() (*mongo.Client, error) {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    log.Println("Connected to MongoDB!")

    return client, nil
}
