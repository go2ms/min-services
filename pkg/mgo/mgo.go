package mgo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"time"
)

var MongoClient *mongo.Client

func Setup() {
	// client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("MONGO connect error: ", err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("mongo Ping errro:", err.Error())
	}

	MongoClient = client
}
