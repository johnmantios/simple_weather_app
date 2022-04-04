package utilities

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	API_KEY string
}

func ParseYaml() string {

	filename, _ := filepath.Abs("./config/api_key.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	return config.API_KEY
}
