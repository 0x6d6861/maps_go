package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDatabase struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDatabase(uri string, database string) *MongoDatabase {

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &MongoDatabase{
		Client:   client,
		Database: client.Database(database),
	}
}

func (database *MongoDatabase) Collection(collection string) *mongo.Collection {
	return database.Database.Collection(collection)
}
