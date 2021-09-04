package configs

import (
	"encoding/json"
	li "library-server/pkg/logger"
	"os"
)

type Config struct {
	DatabaseStrings struct {
		DatabaseName        string `json:"databaseName"`
		DatabaseCollections struct {
			BooksCollectionName string `json:"booksCollectionName"`
		} `json:"databaseCollections"`
	} `json:"databaseStrings"`
	APIPrefix string `json:"apiPrefix"`
	DevEnvs   struct {
		ServerUri    string `json:"serverUri"`
		ServerPort   string `json:"serverPort"`
		WriteTimeout int    `json:"writeTimeout"`
		ReadTimeout  int    `json:"readTimeout"`
		MongodbUri   string `json:"mongodbUri"`
	} `json:"devEnvs"`
	ProdEnvs struct {
	} `json:"prodEnvs"`
	EnvironmentType string `json:"environmentType"`
}

func LoadConfiguration(filePath string, standardLogger *li.StandardLogger) (Config, string, error) {
	var config Config
	configFile, err := os.Open(filePath)
	if err != nil {
		return config, "issue in while loading configuration file", err
	}
	defer configFile.Close()
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return config, "issue in while decoding json from configuration file", err
	}
	return config, "", nil
}
