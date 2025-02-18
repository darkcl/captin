package models

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var cmLogger = log.WithFields(log.Fields{"class": "ConfigurationMapper"})

// ConfigurationMapper - Action to configuration mapper
type ConfigurationMapper struct {
	ActionMap map[string][]Configuration
}

// NewConfigurationMapper - Create ConfigurationMapper with array of Configurations
func NewConfigurationMapper(configs []Configuration) *ConfigurationMapper {
	result := ConfigurationMapper{
		ActionMap: make(map[string][]Configuration),
	}
	for _, config := range configs {
		for _, action := range config.Actions {
			list := result.ActionMap[action]
			list = append(list, config)
			result.ActionMap[action] = list
		}
	}

	return &result
}

// NewConfigurationMapperFromPath - Read Configuration from path
func NewConfigurationMapperFromPath(path string) *ConfigurationMapper {
	pathLogger := cmLogger.WithFields(log.Fields{"path": path})
	data, err := ioutil.ReadFile(path)

	if err != nil {
		pathLogger.Error("Failed to load file")
		panic(err)
	}

	configs := []Configuration{}
	jsonErr := json.Unmarshal(data, &configs)
	if jsonErr != nil {
		pathLogger.Error("Invalid configuration file format")
		panic(jsonErr)
	}

	return NewConfigurationMapper(configs)
}

func (cm ConfigurationMapper) ConfigsForKey(eventKey string) []Configuration {
	return cm.ActionMap[eventKey]
}
