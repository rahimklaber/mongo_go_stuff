package mongo_helper

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	MongoUri = "MONGO_URI"
	DbName   = "orders"
)

func Init() (*mongo.Client, error) {
	uri, ok := os.LookupEnv(MongoUri)
	if !ok {
		return nil, errors.New("env not set")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err // unwrap
	}
	ctx, done := context.WithTimeout(context.Background(), 5*time.Second)
	defer done()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err // unwrap
	}
	initTable(client)
	return client, nil
}

func initTable(client *mongo.Client) {
	databases, err := client.ListDatabases(context.Background(), bson.D{})
	if err != nil {
		log.Fatal("could not list databases")
	}
	found := false
	for i := range databases.Databases {
		if databases.Databases[i].Name == "orders" {
			found = true
			break
		}
	}
	if found {
		log.Println("Database allready created")
		return
	}
	database := client.Database("orders")
	err = database.CreateCollection(context.Background(), DbName)
	if err != nil {
		log.Fatal("could not create database")
	}
}
