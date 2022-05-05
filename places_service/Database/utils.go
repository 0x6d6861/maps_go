package Database

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//var osUri = os.Getenv("MONGODB_URI")

var osUri = "mongodb://localhost:27017/places_cache?authSource=admin"

var DatabaseInstance = NewMongoDatabase(
	osUri,
	"places_cache",
)

var RepositoryInstance = NewRepository(DatabaseInstance)

type TelegramPayload struct {
	Group   string `json:"group"`
	Message string `json:"message"`
}

func SendLogTelegram(message string) {

	url := "https://api.little.bz/internal/telegram/sendMessage"

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	payload := TelegramPayload{
		Group:   "Little_Alerts",
		Message: "SMS_GO_SERVER\n\n" + message,
	}

	requestPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestPayload))

	req.Header.Set("Content-Type", "application/json")

	// req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//Failed to read response.
		// panic(err)
		log.Fatalln(err.Error())
	}

	// TODO: check responses from the provider
	log.Println(string(body))
}
