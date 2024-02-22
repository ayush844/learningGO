package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/ayush844/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// string where our database is located
const connectionString = "mongodb+srv://ayushxyz1625:ayush1625@cluster0.vi4csd3.mongodb.net/?retryWrites=true&w=majority"

// name of our database
const dbName = "netflix"

// name of our collection
const colName = "watchlist"

//most important - taking reference of a mongoDB collection
var collection *mongo.Collection


// CONNECT WITH MONGODB:

//initialization method
func init()  {
	// client option (syntax that we require whenever we make a connection)
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	//Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	client, err := mongo.Connect(context.TODO(), clientOption) 	// TODO returns a non-nil, empty Context.

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongoDB connection success")

	
	collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("collection reference is ready")

}


//MONGODB helpers


// insert 1 record

func insertOneMovie(movie model.Netflix)  { //here model in model.Netflix is package name
	inserted, err := collection.InsertOne(context.Background(), movie);

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted one movie in DB with id: ", inserted.InsertedID);

}

// update 1 movie

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)

}


//delete 1 movie

func deleteOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted, with delete count: ", deleteCount)

}

//delete all movies

func deleteAllMovies() int64{
	filter := bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("all movies got deleted: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}