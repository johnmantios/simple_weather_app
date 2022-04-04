package get_weather

import (
	"io/ioutil"
	"log"
	"net/http"
	"simple_weather_app/utils"
)

var key = utilities.ParseYaml()

func GetWeather(city string) string {

	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + city)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	defer resp.Body.Close()

	//log.Printf(sb)

	return sb

}
