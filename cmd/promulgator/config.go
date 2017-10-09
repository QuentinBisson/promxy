package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ConfigFromFile(path string) (*Config, error) {
	// load the config file
	config := &Config{}
	configBytes, err := ioutil.ReadFile(opts.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %v", err)
	}
	err = yaml.Unmarshal([]byte(configBytes), &config)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling config: %v", err)
	}
	return config, nil
}

// Common configuration for all storage nodes
type Config struct {
	ServerGroups [][]string `yaml:"server_groups"`
}
