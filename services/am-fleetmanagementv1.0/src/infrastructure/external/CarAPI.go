package external

import (
	"encoding/json"
	"fleetmanagement/logic/model"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type CarAPI struct {
	apiURL       string
	CarsEndpoint string
}

type GetCarsApiResponse struct {
	Cars []model.Car `json:"cars"`
}

type GetCarApiResponse struct {
	Vin   model.Vin `json:"vin"`
	Brand string    `json:"brand"`
	Model string    `json:"model"`
}

func NewCarAPI(apiURL string) *CarAPI {
	return &CarAPI{
		apiURL:       apiURL,
		CarsEndpoint: "/cars",
	}
}

func (c *CarAPI) GetCar(vin model.Vin) (model.Car, error) {
	var msg string
	apiURL := fmt.Sprintf("%s%s/%s", c.apiURL, c.CarsEndpoint, vin.Vin)

	// Make the GET request to the API
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		msg = "Failed to make GET request to the Car API"
		log.Error(msg, ": ", err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		msg = fmt.Sprintf("API request failed with status code: %d", resp.StatusCode)
		log.Warn(msg)
		return model.Car{}, fmt.Errorf(msg)
	}

	// Decode the JSON response
	var getCarApiResponse GetCarApiResponse
	err = json.NewDecoder(resp.Body).Decode(&getCarApiResponse)
	if err != nil {
		msg = "Failed to decode JSON response from Car API"
		log.Error(msg, ": ", err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	return model.Car{
		Vin:   getCarApiResponse.Vin,
		Brand: getCarApiResponse.Brand,
		Model: getCarApiResponse.Model,
	}, nil
}
