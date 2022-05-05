package Services

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

type OSMReverseResponse struct {
	Type     string `json:"type"`
	Licence  string `json:"licence"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			PlaceId     int    `json:"place_id"`
			OsmType     string `json:"osm_type"`
			OsmId       int    `json:"osm_id"`
			PlaceRank   int    `json:"place_rank"`
			Category    string `json:"category"`
			Type        string `json:"type"`
			Importance  int    `json:"importance"`
			Addresstype string `json:"addresstype"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
			Address     struct {
				Building    string `json:"building"`
				Road        string `json:"road"`
				Suburb      string `json:"suburb"`
				City        string `json:"city"`
				State       string `json:"state"`
				Postcode    string `json:"postcode"`
				Country     string `json:"country"`
				CountryCode string `json:"country_code"`
			} `json:"address"`
		} `json:"properties"`
		Bbox     []float64 `json:"bbox,omitempty"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry,omitempty"`
	} `json:"features,omitempty"`
}

type OSMConfig struct {
	BaseUrl string
}

type OSM struct {
	BaseService
	Config OSMConfig
}

func NewOSM(config OSMConfig) OSM {
	return OSM{
		BaseService: BaseService{},
		Config:      config,
	}
}

func (service *OSM) GetReverse(query ReverseQuery) (OSMReverseResponse, error) {

	data := OSMReverseResponse{}

	latlng := strings.Split(query.LatLng, ",")

	if len(latlng) != 2 {
		return data, errors.New("invalid format")
	}

	providerPayload := url.Values{}
	providerPayload.Add("format", "geojson")
	providerPayload.Add("lat", latlng[0])
	providerPayload.Add("lon", latlng[1])

	response, err := service.SendGet(providerPayload, service.Config.BaseUrl+"/reverse", RequestConfig{})
	err = json.Unmarshal([]byte(response.Response), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
