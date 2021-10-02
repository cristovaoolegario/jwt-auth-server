package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(env, user, password, hostname, dbname string) (*mongo.Client, error) {
	server := mountServerConnection(env, user, password, hostname, dbname)

	clientOptions := options.Client().ApplyURI(server)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewDatabase(dbName string, client *mongo.Client) *mongo.Database {
	return client.Database(dbName)
}

func mountServerConnection(env, user, password, hostname, dbname string) string {
	if env == "dev" || env == "" {
		return "mongodb://mongo:27017/dev_env"
	}
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, hostname, dbname)
}
