package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sammyhass/weather/controllers"
	"github.com/sammyhass/weather/models"
)

func main() {
	key := os.Getenv("OPEN_WEATHER_API")
	if key == "" {
		fmt.Print("No API Key Found. Please set your 'OPEN_WEATHER_API' environment variable equal to your api key.")
		os.Exit(0)
	}

	apiController := controllers.APIController{
		APIKey: key,
	}

	if len(os.Args[1:]) < 1 {
		// Use 'DEFAULT_CITY_WEATHER_API' ENV Variable
		defaultCity := os.Getenv("DEFAULT_CITY_WEATHER_API")

		if defaultCity == "" {
			fmt.Print("No default city found. Fix this by putting your city in the 'DEFAULT_CITY_WEATHER_API' environment variable")
			os.Exit(0)
		}

		apiController.City = defaultCity

		data, err := apiController.GetWeather()
		if err != nil {
			log.Fatalln(err)
		}
		weather := models.NewWeather(data)

		fmt.Println(weather)

	}
}
