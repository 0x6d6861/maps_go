package Services

import (
	"encoding/json"
	"net/url"
)

type OSMDirectionResponse struct {
	Hints struct {
		VisitedNodesAverage string `json:"visited_nodes.average"`
		VisitedNodesSum     string `json:"visited_nodes.sum"`
	} `json:"hints"`
	Info struct {
		Copyrights []string `json:"copyrights"`
		Took       int      `json:"took"`
	} `json:"info"`
	Paths []struct {
		Distance      float64       `json:"distance"`
		Weight        float64       `json:"weight"`
		Time          int           `json:"time"`
		Transfers     int           `json:"transfers"`
		PointsEncoded bool          `json:"points_encoded"`
		Bbox          []float64     `json:"bbox"`
		Points        string        `json:"points"`
		Legs          []interface{} `json:"legs"`
		Details       struct {
		} `json:"details"`
		Ascend           float64 `json:"ascend"`
		Descend          float64 `json:"descend"`
		SnappedWaypoints string  `json:"snapped_waypoints"`
	} `json:"paths"`
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

func (service *OSM) GetDirection(query DirectionQuery) (OSMDirectionResponse, error) {

	data := OSMDirectionResponse{}

	providerPayload := url.Values{}

	for _, point := range query.Points {
		providerPayload.Add("point", point)
	}

	providerPayload.Add("type", "json")
	providerPayload.Add("vehicle", "car")
	providerPayload.Add("weighting", "fastest")
	providerPayload.Add("elevation", "false")
	providerPayload.Add("instructions", "false")
	providerPayload.Add("key", "d922f342-53fc-4791-9fdc-5f2a39522c71")
	providerPayload.Add("locale", "en-US")

	response, err := service.SendGet(providerPayload, service.Config.BaseUrl+"/route", RequestConfig{})
	err = json.Unmarshal([]byte(response.Response), &data)

	// fmt.Println(response.Response)

	if err != nil {
		return data, err
	}

	return data, nil
}
