package cms_db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_cms_db() (*mongo.Client, error) {
	err := godotenv.Load()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI environment variables must be set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Disconnect_cms_db(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
