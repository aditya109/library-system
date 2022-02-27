package config

import (
	"log"
	"os"

	"github.com/aditya109/library-system/internal/models"
	"github.com/aditya109/library-system/pkg/helper"
	"gopkg.in/yaml.v2"
)

// LoadConfiguration retrieves configuration from config file
func LoadConfiguration() (*models.Config, error) {
	var config models.Config
	var configFilePath string = helper.GetAbsolutePath("/config/config.yaml")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &config, nil
}
