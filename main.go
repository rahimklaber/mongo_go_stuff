package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongohelper "goproject/mongo-helper"
	"log"
)

type Comment struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	MovieId primitive.ObjectID `bson:"movie_id,omitempty"`
	Text    string             `bson:"text,omitempty"`
	Date    primitive.DateTime `bson:"date,omitempty"`
}

func main() {

	/*
	   Connect to my cluster
	*/
	client, err := mongohelper.Init()
	ctx := context.Background()
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
