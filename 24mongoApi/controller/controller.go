package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ayushxyz1625:ayush1625@cluster0.vi4csd3.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"

const colName = "watchlist"

//most important
var collection *mongo.Collection


// connect with mongoDB

func init()  {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("collection reference is ready")

}
