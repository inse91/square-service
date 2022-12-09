package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDbUrl string = "mongodb://localhost:27017"
)

func NewMongoClient(ctx context.Context, database string) (db *mongo.Database, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	clientOpts := options.Client()
	clientOpts.ApplyURI(mongoDbUrl)
	clientOpts.SetAuth(options.Credential{
		AuthSource: database,
		Username:   "",
		Password:   "",
	})

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	db = client.Database(database)
	return
}
