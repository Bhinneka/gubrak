package gubrak

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config struct
type Config struct {
	BaseURL string            `json:"baseUrl"`
	Headers map[string]string `json:"headers"`
}

// LoadConfig for load config from json file
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")

	if err != nil {
		return nil, ErrorConfigNotFound
	}

	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(jsonData, &config)

	if err != nil {
		return nil, ErrorFormatConfigNotValid
	}

	return &config, nil
}
