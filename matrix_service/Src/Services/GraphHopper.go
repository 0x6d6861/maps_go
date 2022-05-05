package Services

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
)

type GraphHopperMatrixQueryResponse struct {
	Hints struct {
		VisitedNodesAverage string `json:"visited_nodes.average"`
		VisitedNodesSum     string `json:"visited_nodes.sum"`
	} `json:"hints"`
	Info struct {
		Copyrights []string `json:"copyrights"`
		Took       int      `json:"took"`
	} `json:"info"`
	Paths []struct {
		Distance         float64 `json:"distance"`
		Weight           float64 `json:"weight"`
		Time             int     `json:"time"`
		Transfers        int     `json:"transfers"`
		SnappedWaypoints string  `json:"snapped_waypoints"`
	} `json:"paths"`
}

type GraphHopperConfig struct {
	BaseUrl string
}

type GraphHopper struct {
	BaseService
	Config GraphHopperConfig
}

func NewGraphHopper(config GraphHopperConfig) GraphHopper {
	return GraphHopper{
		BaseService: BaseService{},
		Config:      config,
	}
}

func (service *GraphHopper) GetMatrix(payload MatrixQuery) (MatrixQueryResponse, error) {

	var serviceResponse = MatrixQueryResponse{
		DestinationAddresses: payload.Destinations,
		OriginAddresses:      payload.Origins,
		Status:               "OK",
	}

	ch := make(chan MatrixQueryRowElementResponse, len(payload.Destinations)*len(payload.Origins))

	for _, origin := range payload.Origins {

		for _, destination := range payload.Destinations {
			go func(_origin string, _destination string) {
				providerPayload := url.Values{}

				providerPayload.Add("point", _origin)
				providerPayload.Add("point", _destination)

				providerPayload.Add("vehicle", "car")
				providerPayload.Add("debug", "false")
				providerPayload.Add("key", "d922f342-53fc-4791-9fdc-5f2a39522c71")
				providerPayload.Add("type", "json")
				providerPayload.Add("instructions", "false")
				providerPayload.Add("calc_points", "false")
				providerPayload.Add("points_encoded", "true")

				response, err := service.SendGet(providerPayload, service.Config.BaseUrl+"/route", RequestConfig{})
				if err != nil {
					fmt.Println(err.Error())
				}
				data := GraphHopperMatrixQueryResponse{}
				err = json.Unmarshal([]byte(response.Response), &data)
				if err != nil {
					fmt.Println(err.Error())
				}
				var serviceElementResponse = MatrixQueryElementResponse{
					Distance: struct {
						Text  string `json:"text"`
						Value int    `json:"value"`
					}{
						Text:  strconv.FormatFloat(data.Paths[0].Distance/1000, 'f', -1, 32) + " Km",
						Value: int(math.Round(data.Paths[0].Distance)),
					},
					Duration: struct {
						Text  string `json:"text"`
						Value int    `json:"value"`
					}{
						Text:  strconv.FormatFloat(float64(data.Paths[0].Time/60000), 'f', -1, 32) + " mins",
						Value: int(math.Round(float64(data.Paths[0].Time / 1000))),
					},
					Status: "OK",
				}

				var serviceRowElementResponse = MatrixQueryRowElementResponse{}
				serviceRowElementResponse.Elements = append(serviceRowElementResponse.Elements, serviceElementResponse)

				ch <- serviceRowElementResponse

			}(origin, destination)
		}

	}

	for i := 0; i < len(payload.Origins)*len(payload.Destinations); i++ {
		respChan := <-ch
		serviceResponse.Rows = append(serviceResponse.Rows, respChan)
	}

	return serviceResponse, nil
}
