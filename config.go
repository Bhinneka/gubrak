package gubrak

import (
	"encoding/json"
	"io/ioutil"
)

// Config struct
type Config struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Payload interface{}       `json:"payload"`
}

// LoadConfig for load config from json file
func LoadConfig(source string) (*Config, error) {
	jsonData, err := ioutil.ReadFile(source)

	if err != nil {
		return nil, ErrorConfigNotFound
	}

	var config Config

	err = json.Unmarshal(jsonData, &config)

	if err != nil {
		return nil, ErrorFormatConfigNotValid
	}

	return &config, nil
}
