package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoSocket struct {
	DbScocket *mongo.Client
	DbAddr    string
	Ctx       *context.Context
}

func Init(address string, ctx *context.Context) *MongoSocket {

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Connect(*ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected.")
	}
	//defer func() { mongoClient.Disconnect(*ctx); fmt.Println("database disconnected") }()

	return &MongoSocket{Ctx: ctx, DbAddr: address, DbScocket: mongoClient}
}

// GetCollection will return collection and database
func (socket *MongoSocket) GetCollection(dbName string, collectionName string) (*mongo.Collection, *mongo.Database) {
	db := socket.DbScocket.Database(dbName)
	collection := db.Collection(collectionName)

	return collection, db
}

// InsertData Insert a doc with a filter
func (socket *MongoSocket) InsertData(collection *mongo.Collection, filter interface{}, doc interface{}) {
	count, err := collection.CountDocuments(*socket.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		insertResult, err := collection.InsertOne(*socket.Ctx, doc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("inserted one document", insertResult.InsertedID)
	} else {
		fmt.Println("the document already exists")
	}
}

func (socket *MongoSocket) FetchAll(sports *[]bson.M, odds *[]bson.M) int {
	SportsCollection, _ := socket.GetCollection("sports", "gp3")
	OddsCollection, _ := socket.GetCollection("sports", "gp4")

	cursor, err := SportsCollection.Find(*socket.Ctx, bson.M{})
	if err != nil {
		return 1
	}
	if err = cursor.All(*socket.Ctx, sports); err != nil {
		log.Fatal(err)
	}

	cursor1, err := OddsCollection.Find(*socket.Ctx, bson.M{})
	if err != nil {
		return 1
	}
	if err = cursor1.All(*socket.Ctx, odds); err != nil {
		log.Fatal(err)
	}

	return 0
}
