package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"square-service/internal/config"
)

// func NewMongoClient(ctx context.Context, host, port, database string) (db *mongo.Database, err error) {

func NewMongoClient(ctx context.Context, cfg config.MongoDB) (db *mongo.Database, err error) {

	//url := fmt.Sprintf("mongodb://%s:%s@%s:%s", login, password, host, port)
	url := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Password)
	clientOpts := options.Client().ApplyURI(url)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf(`error while trying to connect to db: %v`, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(`error while trying to ping db: %v`, err)
	}

	db = client.Database(cfg.DataBase)
	return
}
