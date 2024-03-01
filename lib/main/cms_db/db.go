package cms_db

import (
	"context"
	"fmt"
	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect_cms_db() (*mongo.Client, error) {
	// uri := os.Getenv("MONGODB_URI")
	uri := "mongodb+srv://bernarmuller:@C8PF7RngnA2CX*@portal.eooko63.mongodb.net/?retryWrites=true&w=majority&appName=portal"
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI environment variables must be set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func disconnect_cms_db(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
