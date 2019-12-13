package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sammyhass/weather/models"
)

// APIUrl is the base url required to retrieve the weather from the api
var APIUrl = "http://api.openweathermap.org/data/2.5/weather?"

// APIController contains data provided by the user including City and ApiKey
type APIController struct {
	APIKey string
	City   string
}

// GetWeather retrieves the current weather and returns a models.ApiResponse
func (c *APIController) GetWeather() (models.ApiResponse, error) {
	if c.City == "" {
		return models.ApiResponse{}, fmt.Errorf("no city set")
	}

	res, err := http.Get(fmt.Sprintf("%sAPPID=%s&q=%s", APIUrl, c.APIKey, c.City))
	if err != nil {
		return models.ApiResponse{}, err
	}
	var jsonRes models.ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		return models.ApiResponse{}, err
	}

	if jsonRes.Cod != 200 {
		return models.ApiResponse{}, fmt.Errorf("something went wrong! Check your API Key is valid")
	}

	return jsonRes, nil
}
