package Services

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type BaseService struct {
}

type ServiceProviderResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Response  string `json:"response,omitempty"`
	Responses []ServiceProviderResponse
}

type RequestConfig struct {
	Auth struct {
		Username string
		Password string
	}
	Headers map[string]string
}

type ReverseResponseResponse struct {
}

type ReverseQuery struct {
	LatLng string `form:"latlng" json:"latlng" uri:"latlng" binding:"required"`
}

func (provider *BaseService) SendGet(payload url.Values, url string, config RequestConfig) (serviceResponse *ServiceProviderResponse, err error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.URL.RawQuery = payload.Encode()

	// TODO: refactor this
	if config.Auth.Username != "" {
		req.SetBasicAuth(config.Auth.Username, config.Auth.Password)
	}

	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	// req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		//panic(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//Failed to read response.
		// log.Fatalln(err.Error())
		return nil, err
	}

	if response.StatusCode >= http.StatusOK && response.StatusCode < 300 {
		//var jsonStr ServiceResponse
		//Convert bytes to String and print
		//json.Unmarshal(body, &jsonStr)
		// TODO: check responses from the provider
		return &ServiceProviderResponse{
			Success:  true,
			Message:  "Success",
			Response: string(body),
		}, nil

	} else {
		//The status is not Created. print the error.
		return &ServiceProviderResponse{
			Success:  false,
			Message:  "Error",
			Response: string(body),
		}, nil
	}
}

func (provider *BaseService) SendPost(payload []byte, url string, config RequestConfig) (serviceResponse *ServiceProviderResponse, err error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	// TODO: refactor this
	if config.Auth.Username != "" {
		req.SetBasicAuth(config.Auth.Username, config.Auth.Password)
	}

	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	// req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		//panic(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//Failed to read response.
		// panic(err)
		log.Fatalln(err.Error())
		return nil, err
	}
	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		//var jsonStr ServiceResponse
		//Convert bytes to String and print
		//json.Unmarshal(body, &jsonStr)
		// TODO: check responses from the provider
		return &ServiceProviderResponse{
			Success:  true,
			Message:  "Success",
			Response: string(body),
		}, nil

	} else {
		//The status is not Created. print the error.
		return &ServiceProviderResponse{
			Success:  false,
			Message:  "Error occurred",
			Response: string(body),
		}, nil
	}
}
