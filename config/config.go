package config

import (
	"encoding/json"
	"os"

	logger "github.com/ryoccd/gochat/log"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// config
var Config Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		logger.Error("Cannot open config file", err)
	}

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		logger.Error("Cannot get configration from file", err)
	}

}

// func GetConfig(propName string) interface{} {
// 	switch propName {
// 	case "address":
// 		return conf.Address
// 	case "readTimeout":
// 		return conf.ReadTimeout
// 	case "writeTimeout":
// 		return conf.WriteTimeout
// 	case "static":
// 		return conf.Static
// 	default:
// 		logger.Error("wrong propName", propName)
// 		return nil
// 	}
// }

func Version() string {
	return "v0.0.1"
}
