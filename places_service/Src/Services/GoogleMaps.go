package Services

import (
	"context"
	"googlemaps.github.io/maps"
	"strconv"
	"strings"
)

type ProviderConfig struct {
	Key string
}

type GoogleMaps struct {
	BaseService
	Client maps.Client
	Config ProviderConfig
}

func NewGoogleMaps(config ProviderConfig) GoogleMaps {
	c, err := maps.NewClient(maps.WithAPIKey(config.Key))
	if err != nil {
		panic(err)
	}

	return GoogleMaps{
		BaseService: BaseService{},
		Config:      config,
		Client:      *c,
	}
}

func (service *GoogleMaps) AutoComplete(query PlaceQuery) ([]PlaceQueryResponse, error) {
	latlng := strings.Split(query.Location, ",")
	lat, _ := strconv.ParseFloat(latlng[0], 64)
	lng, _ := strconv.ParseFloat(latlng[1], 64)
	location := maps.LatLng{
		Lat: lat,
		Lng: lng,
	}
	r := &maps.PlaceAutocompleteRequest{
		Input:        query.Q,
		Location:     &location,
		Radius:       500000,
		StrictBounds: false,
		SessionToken: maps.PlaceAutocompleteSessionToken{},
	}
	autocomplete, err := service.Client.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		return nil, err
	}

	var predictions = []PlaceQueryResponse{}

	for _, prediction := range autocomplete.Predictions {
		predictions = append(predictions, PlaceQueryResponse{
			Id:          prediction.PlaceID,
			Description: prediction.StructuredFormatting.MainText,
			Street:      prediction.StructuredFormatting.SecondaryText,
			State:       prediction.StructuredFormatting.SecondaryText,
		})
	}

	return predictions, nil
}
