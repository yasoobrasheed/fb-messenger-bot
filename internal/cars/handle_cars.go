package cars

import (
	"encoding/json"
	"fbmessenger_bot/util"
	"io"
	"log"
	"net/http"
)

type CarsRequest struct {
	SenderId string `json:"sender_id"`
}

type CarsResponse struct {
	Id           string `json:"id"`
	Image        string `json:"image"`
	Vehicle      string `json:"vehicle"`
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"productName"`
	CreatedAt    string `json:"createdAt"`
}

func HandleCars(w http.ResponseWriter, r *http.Request) {
	var carsRequest CarsRequest

	err := util.ParseAndUnmarshallRequestBody(r, &carsRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := MakeCarsApiRequest()
	if err != nil {
		log.Println("error making cars api request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carsResponse []CarsResponse

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading res body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(responseBody, &carsResponse)
	if err != nil {
		log.Println("error unmarshalling res body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	log.Println(carsResponse)

	w.WriteHeader(http.StatusOK)
}

func MakeCarsApiRequest() (*http.Response, error) {
	URL := "https://62daf70dd1d97b9e0c49ca5d.mockapi.io/v1/products"

	response, err := http.Get(URL)
	if err != nil {
		log.Println("Error posting message via FB Graph API:", err)
		return nil, err
	}

	return response, nil
}
