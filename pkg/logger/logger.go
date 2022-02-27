package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/aditya109/library-system/internal/models"
	cfg "github.com/aditya109/library-system/pkg/config"
	log "github.com/sirupsen/logrus"
)

var (
	config *models.Config
	err    error
)

// InitializeLogging returns a configured logger object
func InitializeLogging() {
	// initializing a configuration object
	config, err = cfg.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
		return 
	}
	
}


// ConfigureLogger configures logger based on logging configuration present in config.json
func ConfigureLogger(loggingConfiguration *models.LevelledLogs) error {
	// declaring writes to stores all the enabled io writers
	var writes []io.Writer
	if loggingConfiguration.EnableLoggingToFile { // if enabled, configures logger to log into file at location of persistence
		logFileName, err := her

	}
}