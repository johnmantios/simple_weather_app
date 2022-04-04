package main

import (
	"fmt"
	"simple_weather_app/get_weather"
)

func main() {
	message := get_weather.GetWeather("Thessaloniki")
	fmt.Println(message)
}
