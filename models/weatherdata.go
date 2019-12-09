package models

import "github.com/sammyhass/weather/helpers/temperature"

type Weather struct {
	Location Location
	Info WeatherInfo
}

type WeatherInfo struct {
	TempF temperature.Temp
	Desc string
}

type Location struct {
	Name string
	Lat float64
	Lon float64
}

func NewWeather(response ApiResponse) *Weather {
	w := new(Weather)
	w.Info.Desc = response.Weather[0].Main
	w.Info.TempF = temperature.Temp{
		Val: response.Main.Temp,
		Unit: temperature.K,
	}.To(temperature.F)
	
	w.Location = Location{
		Name: response.Name,
		Lat:  response.Coord.Lat,
		Lon:  response.Coord.Lon,
	}

	return w
}