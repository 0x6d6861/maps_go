package Services

import (
	"mpasGo/Database"
	"net/url"
	"strings"
)

type MongoQueryResponse struct {
}

type MongoConfig struct {
	RepositoryInstance *Database.Repository
}

type Mongo struct {
	BaseService
	Config             MongoConfig
	RepositoryInstance *Database.Repository
}

func NewMongo(config MongoConfig) Mongo {
	return Mongo{
		BaseService:        BaseService{},
		Config:             config,
		RepositoryInstance: config.RepositoryInstance,
	}
}

func (service *Mongo) AutoComplete(query PlaceQuery) ([]PlaceQueryResponse, error) {
	latlng := strings.Split(query.Location, ",")

	providerPayload := url.Values{}
	providerPayload.Add("lat", latlng[0])
	providerPayload.Add("lon", latlng[1])
	providerPayload.Add("location_bias_scale", "5")
	providerPayload.Add("q", query.Q)

	results, err := service.RepositoryInstance.SearchDBPlace(Database.SearchDBPlaceQuery{
		Query: query.Q,
		//Days:     0,
		Country:  query.Country,
		City:     query.City,
		Location: query.Location,
	})

	if err != nil {
		return nil, err
	}

	var predictions = []PlaceQueryResponse{}

	for _, result := range results {
		predictions = append(predictions, PlaceQueryResponse{
			Id:          result.PlaceId,
			Description: result.Description,
			Street:      result.Street,
			LatLng:      result.LatLng,
			City:        result.City,
			Country:     result.Country,
		})
	}

	return predictions, nil
}
