package main

import (
	"os"
	"simple_weather_app/get_weather"
)

func main() {
	number_of_args := 4 //TODO: change to arbitrary number of arguments
	for argument := 1; argument < number_of_args; argument++ {
		get_weather.GetWeather(os.Args[argument])
	}

}
