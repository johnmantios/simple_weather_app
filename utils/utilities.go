package utilities

import (
	"github.com/goccy/go-yaml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	API_KEY string
}

func ParseYaml() string {

	filename, _ := filepath.Abs("./config/api_key.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return config.API_KEY
}

func Logger(my_string string) {
	// log to custom file
	LOG_FILE := "./simple_weather_app_log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	defer logFile.Close()

	//Setting a multiwriter to write to both a log file and to saved stdout
	mw := io.MultiWriter(os.Stdout, logFile)

	//Set log output. Writes with log.Print will also write to mw
	log.SetOutput(mw)

	//log date-time, filename and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(my_string)

}

func FloatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}
