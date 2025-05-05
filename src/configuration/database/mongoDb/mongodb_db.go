package mongoDb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbURI := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_USER_DB)

	// CLIENT
	client, err := mongo.Connect(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	// PING
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbDatabase), nil
}
