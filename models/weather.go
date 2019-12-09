package models

import (
	"fmt"
	"time"
)

type Weather struct {
	Location Location
	Info WeatherInfo
}

type WeatherInfo struct {
	Temp Temp
	Desc  string
	Time time.Time
}

type Location struct {
	Name string
	Lat float64
	Lon float64
}

func NewWeather(response ApiResponse) *Weather {
	w := new(Weather)
	w.Info.Desc = response.Weather[0].Main
	w.Info.Time = time.Now()
	w.Info.Temp = Temp{
		Val:  response.Main.Temp,
		Unit: K,
	}.To(F)
	
	w.Location = Location{
		Name: response.Name,
		Lat:  response.Coord.Lat,
		Lon:  response.Coord.Lon,
	}

	return w
}

func (w Weather) String() string {
	t := w.Info.Time
	format := t.Format(time.Kitchen)
	return fmt.Sprintf("%s %s : %v, %s", format, w.Location.Name, w.Info.Temp, w.Info.Desc)
}