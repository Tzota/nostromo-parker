package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tzota/nostromo-parker/internal/kind"
)

// Point represent a sensor in config file
type Point struct {
	Mac  string    `json:"mac"`
	Kind kind.Kind `json:"kind"`
	Skip bool      `json:"skip"`
}

// Config is a general entry point
type Config struct {
	Points []Point `json:"points"`
}

// ReadFromFile your config
func ReadFromFile(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	return read(data)
}

// Read you config from bytes
func read(data []byte) (Config, error) {
	cfg := Config{}
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
