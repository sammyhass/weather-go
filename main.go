package main

import (
	"fmt"
	"github.com/sammyhass/weather/controllers"
	"log"
	"os"
)

func main() {
	key := os.Getenv("OPEN_WEATHER_API")
	if key == "" {
		fmt.Print("No API Key Found. Please set your 'OPEN_WEATHER_API' environment variable equal to your api key.")
		os.Exit(0)
	}

	apiController := controllers.ApiController{
		ApiKey:key,
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

		fmt.Printf("Weather - %s, %s\n", apiController.City, data.Sys.Country)

		fmt.Printf("%.1f F and %s\n", data.Main.Temp * 9 /5 - 459.67, data.Weather[0].Main)

	}
}
