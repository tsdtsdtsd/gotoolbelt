// Package config provides structures for application config
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// A Config represents the complete application configuration
type Config struct {
	HTTPD   HTTPD   `yaml:"httpd"`
	Storage Storage `yaml:"storage"`
	Worker  Worker  `yaml:"worker"`
}

// HTTPD is the HTTP daemon related sub-config
type HTTPD struct {
	Address       string `yaml:"address"`
	ImagesBaseURI string `yaml:"imagesBaseURI"`
}

// Storage is the storage related sub-config
type Storage struct {
	MongoDB       MongoDB `yaml:"mongodb"`
	Folders       Folders `yaml:"folders"`
	FallbackImage string  `yaml:"fallbackImage"`
}

// Worker is the worker related sub-config
type Worker struct {
	Quantity uint16 `yaml:"quantity"`
}

// MongoDB is the MongoDB related sub-config
type MongoDB struct {
	Server    string `yaml:"server"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
	Timeout   int16  `yaml:"timeout"`
	Source    string `yaml:"source"`
	Mechanism string `yaml:"mechanism"`
}

// Folders contains storage folders on disk
type Folders struct {
	Master    string `yaml:"master"`
	Generated string `yaml:"generated"`
}

// NewFromYAMLFile returns a new Config struct, filled with values from given yaml filepath
func NewFromYAMLFile(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
