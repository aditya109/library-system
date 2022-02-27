package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aditya109/library-system/internal/models"
	"github.com/aditya109/library-system/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

// GetConfiguration retrieves configuration from config file or environment configuration(TODO:)
func GetConfiguration() (*models.Config, error) {
	// declaring a config object
	var config = models.Config{}
	// getting the absolute file path of the config file
	var configFilePath, err = helper.GetAbsolutePath("/config/config.json")
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}

	// storing the content of config file in a config object
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}
	return &config, nil
}
