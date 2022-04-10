package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"truffle.io/storage"
)

type Config struct {
	DBOptions *storage.NewDatabaseOptions `json:"database"`
}

func GetConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
