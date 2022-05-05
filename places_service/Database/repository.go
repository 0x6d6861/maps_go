package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Repository struct {
	MongoDatabase *MongoDatabase
}

type SearchDBPlaceQuery struct {
	Query    string `json:"query,omitempty" bson:"query"`
	Days     int    `json:"days,omitempty" bson:"days,omitempty"`
	Country  string `json:"country,omitempty" bson:"country,omitempty"`
	City     string `json:"city,omitempty" bson:"city,omitempty"`
	Location string `json:"location,omitempty" bson:"location,omitempty"`
}

type PlaceModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	PlaceId     string             `bson:"placeId" json:"PlaceId"`
	Description string             `bson:"description" json:"description"`
	Street      string             `bson:"street" json:"street"`
	LatLng      string             `bson:"latlng" json:"latlng"`
	City        string             `bson:"city" json:"city"`
	Country     string             `bson:"country" json:"country"`
	ExpiresAt   primitive.DateTime `bson:"expiresAt" json:"expiresAt"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
}

func NewRepository(mongoDatabase *MongoDatabase) *Repository {
	return &Repository{MongoDatabase: mongoDatabase}
}
func (r Repository) InsertPlaces(places []PlaceModel) (*mongo.InsertManyResult, error) {

	var toInsert []interface{}

	for _, place := range places {
		place.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		toInsert = append(toInsert, place)
	}
	result, err := r.MongoDatabase.Collection("placemodels").InsertMany(context.TODO(), toInsert)
	return result, err
}

// TODO: add method to find place by city, latlng
func (r Repository) SearchDBPlace(searchDBPlaceQuery SearchDBPlaceQuery) ([]*PlaceModel, error) {
	opts := options.Find()
	opts.SetLimit(10)

	var allResults []*PlaceModel

	var dbQuery = bson.D{}

	if searchDBPlaceQuery.City != "" {
		dbQuery = append(dbQuery, bson.E{"city", bson.D{
			{"$regex", primitive.Regex{
				Pattern: "^" + searchDBPlaceQuery.City + "$",
				Options: "i",
			},
			},
		},
		})
	}

	if searchDBPlaceQuery.City != "" {
		dbQuery = append(dbQuery, bson.E{"country", bson.D{
			{"$regex", primitive.Regex{
				Pattern: "^" + searchDBPlaceQuery.Country + "$",
				Options: "i",
			},
			},
		},
		})
	}

	dbQuery = append(dbQuery, bson.E{"description", bson.D{
		{"$regex", primitive.Regex{
			Pattern: searchDBPlaceQuery.Query,
			Options: "i",
		},
		},
	},
	})

	results, err := r.MongoDatabase.Collection("placemodels").Find(context.TODO(), dbQuery, opts)

	for results.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem PlaceModel
		err := results.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		allResults = append(allResults, &elem)
	}

	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	return allResults, err
}
