package config

import (
	"encoding/json"
	"os"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// config
var Conf Configuration

func init() {
	loadConfig()
}

//Reads the configuration file and converts it to a format that can be read in the project.
func loadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		logger.Error("Cannot open config file", err)
	}

	logger.Info("Success! Open config file")
	decoder := json.NewDecoder(file)
	Conf = Configuration{}
	err = decoder.Decode(&Conf)
	if err != nil {
		logger.Error("Cannot get configration from file", err)
	}

	logger.Info(Conf)
	logger.Info("End! Load config file")
}

func Version() string {
	return "v0.0.1"
}
