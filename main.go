package main

import (
	"simple_weather_app/get_weather"
	"simple_weather_app/utils"
)

func main() {
	message := get_weather.GetWeather("Thessaloniki")
	utilities.Logger(message)
}
