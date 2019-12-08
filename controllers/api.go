package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/sammyhass/weather/models"
)

var ApiUrl = "http://api.openweathermap.org/data/2.5/weather?"
type ApiController struct {
	ApiKey string
	City string
}


func (c *ApiController) GetWeather() (models.ApiResponse, error) {
	if c.City == "" {
		return models.ApiResponse{}, fmt.Errorf("no city set")
	}

	res, err := http.Get(fmt.Sprintf("%sAPPID=%s&q=%s", ApiUrl, c.ApiKey, c.City))
	if err != nil {
		return models.ApiResponse{}, err
	}
	var jsonRes models.ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		return models.ApiResponse{}, err
	}

	if jsonRes.Cod != 200 {
		return models.ApiResponse{}, fmt.Errorf("Something went wrong! Check your API Key is valid.")
	}

	return jsonRes, nil
}