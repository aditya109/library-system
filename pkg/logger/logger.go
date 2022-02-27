package logwrapper

import (
	"io"
	"os"

	"github.com/aditya109/library-system/internal/models"
	"github.com/aditya109/library-system/pkg/config"
	"github.com/aditya109/library-system/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

// InitializeLogging returns a configured logger object
func InitializeLogging() error {
	// initializing a configuration object
	config, err := config.GetConfiguration()
	if err != nil {
		logger.Error(err)
		return err
	}
	if err := ConfigureLogger(config.ApplicationConfig.LevelledLogs); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// ConfigureLogger configures logger based on logging configuration present in config.*.json
func ConfigureLogger(loggingConfig models.LevelledLogs) error {
	// declaring writers to location models.PersistenceLocationtore all the enabled the io writers
	var writers []io.Writer
	if loggingConfig.EnableLoggingToFile { // if enabled, configuring logger to log into file by filespecs of logging configuration
		logFileName, err := helper.GetFormattedFileName(loggingConfig.PersistenceLocation)
		if err != nil {
			logger.Errorf("error in getting filename: %v", err)
			return err
		}
		f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			logger.Errorf("error opening file: %v", err)
			return err
		}
		writers = append(writers, f)
	}

	if loggingConfig.EnableLoggingToStdout { // if enabled, configuring logger to log into stdout, according to logging configuration
		writers = append(writers, os.Stdout)
	}

	mw := io.MultiWriter(writers...)
	logger.SetOutput(mw)

	if loggingConfig.OutputFormatter == "json" { // configuring log-syntax type format - json/text
		logger.SetFormatter(&logger.JSONFormatter{})
	} else {
		logger.SetFormatter(&logger.TextFormatter{
			DisableColors: !loggingConfig.EnableColors,
			FullTimestamp: loggingConfig.EnableFullTimestamp,
		})
	}
	return nil
}
