package get_weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	var data map[string]interface{}

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data["current"])

}
