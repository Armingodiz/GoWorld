package main

import (
	"context"
	"fmt"
	//"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb+srv://armin:<armin3011>@cluster0.v7eej.mongodb.net/<test>?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("tsss1")
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println("tsss2")
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("tsssss3")
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
}
