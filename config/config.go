package config

import (
    "context"
    "fmt"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/option"
)

func ConnectToMongoDB() (*mongo.Client, error) {
    uri := os.Getenv("MONGODB_URI)
    if uri == "" {
        return nil, fmt.Errorf("MONGODB_URI is not set")
    }
    
    clientOption := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.Background(), clientOption)
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    return client, nil
}
