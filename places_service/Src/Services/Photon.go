package Services

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type PhotonQueryResponse struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Type       string `json:"type"`
		Properties struct {
			OsmId       int       `json:"osm_id"`
			OsmType     string    `json:"osm_type"`
			Extent      []float64 `json:"extent"`
			Country     string    `json:"country"`
			OsmKey      string    `json:"osm_key"`
			City        string    `json:"city"`
			Street      string    `json:"street"`
			Countrycode string    `json:"countrycode"`
			OsmValue    string    `json:"osm_value"`
			Postcode    string    `json:"postcode"`
			Name        string    `json:"name"`
			State       string    `json:"state"`
		} `json:"properties"`
	} `json:"features"`
	Type string `json:"type"`
}

type PhotonConfig struct {
	ProviderConfig
	BaseUrl string
}

type Photon struct {
	BaseService
	Config PhotonConfig
}

func NewPhoton(config PhotonConfig) Photon {
	return Photon{
		BaseService: BaseService{},
		Config:      config,
	}
}

func (service *Photon) AutoComplete(query PlaceQuery) ([]PlaceQueryResponse, error) {
	latlng := strings.Split(query.Location, ",")

	providerPayload := url.Values{}
	providerPayload.Add("lat", latlng[0])
	providerPayload.Add("lon", latlng[1])
	providerPayload.Add("location_bias_scale", "5")
	providerPayload.Add("q", query.Q)

	response, err := service.SendGet(providerPayload, service.Config.BaseUrl+"/api", RequestConfig{})
	if err != nil {
		return nil, err
	}

	data := PhotonQueryResponse{}
	err = json.Unmarshal([]byte(response.Response), &data)
	if err != nil {
		return nil, err
	}

	var predictions = []PlaceQueryResponse{}

	for _, feature := range data.Features {
		condition := false
		if strings.Contains(strings.ToLower(feature.Properties.Name), strings.ToLower(strings.Split(query.Q, " ")[0])) {

			if query.City != "" && query.Country != "" {
				condition = strings.ToLower(query.City) == strings.ToLower(feature.Properties.City) && strings.ToLower(query.Country) == strings.ToLower(feature.Properties.Country)
			}

			if query.City != "" || query.Country != "" {
				condition = strings.ToLower(query.City) == strings.ToLower(feature.Properties.City) || strings.ToLower(query.Country) == strings.ToLower(feature.Properties.Country)
			}

			if condition {
				predictions = append(predictions, PlaceQueryResponse{
					Id:          strconv.Itoa(feature.Properties.OsmId),
					Description: feature.Properties.Name,
					Street:      feature.Properties.Street,
					State:       feature.Properties.State,
					LatLng:      strconv.FormatFloat(feature.Geometry.Coordinates[1], 'f', -1, 64) + "," + strconv.FormatFloat(feature.Geometry.Coordinates[0], 'f', -1, 64),
					City:        feature.Properties.Country,
					Country:     feature.Properties.Country + ".",
				})
			}
		}
	}

	return predictions, nil
}
