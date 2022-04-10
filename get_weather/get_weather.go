package get_weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple_weather_app/api_model"
	"simple_weather_app/utils"
)

var key = utilities.ParseYaml()

func GetWeather(city string) {

	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + city)
	if err != nil {
		utilities.Logger("Something went wrong whilst getting the api response")
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utilities.Logger("An error ocurred while trying to read the api response body")
		panic(err)
	}
	defer resp.Body.Close()

	res := &api_model.CurrentJsonResponse{}

	err = json.Unmarshal([]byte(body), res)
	if err != nil {
		panic(err)
	}

	fmt.Println("The weather in " + city + " at " + *(res.Current.LastUpdated) + " was: ")
	fmt.Println(*(res.Current.TempC))

}
