// Package sample http API
//
// Documentation of crud-template API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aditya109/library-system/internal/constants"
	"github.com/aditya109/library-system/internal/models"
	rt "github.com/aditya109/library-system/internal/router"
	cfg "github.com/aditya109/library-system/pkg/config"
	logger "github.com/sirupsen/logrus"
)

var (
	config       *models.Config
	httpPort     string
	prefix       string
	endpoint     string
	writeTimeout time.Duration
	readTimeout  time.Duration
	err          error
	envs         models.Envs
)

func Start() {
	config, err = cfg.GetConfiguration() // retrieving configuration
	if err != nil {
		logger.Fatal(err)
		return
	}

	getApplicableEnvironmentVariablesFromConfig() // getting applicable environment variables from config
	setHTTPPortFromConfigObject()                 // getting http port from config
	setEndpointFromConfigObject()                 // getting endpoint from config
	setTimeoutsFromConfigObject()                 // getting timeouts from config

	// configuring router for the server
	router := rt.ConfigureRouter()
	logger.Info("router configuration successful")
	logger.Info(fmt.Sprintf("starting server at %s://%s", prefix, endpoint))
	logger.Info(fmt.Sprintf("swagger docs can be viewed at %s://%s/docs", prefix, endpoint))

	// configuring server
	srv := &http.Server{
		Handler:      router,
		Addr:         endpoint,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	logger.Fatal(srv.ListenAndServe())
}

// getApplicableEnvironmentVariablesFromConfig gets applicable environment variables from configuration
func getApplicableEnvironmentVariablesFromConfig() {
	switch config.ServerConfig.EnvironmentType {
	case constants.DEV:
		envs = config.ServerConfig.DevEnvs
	case constants.STAGING:
		envs = config.ServerConfig.StagEnvs
	case constants.PRODUCTION:
		envs = config.ServerConfig.ProdEnvs
	default:
		envs = config.ServerConfig.DevEnvs
	}
}

// setHTTPPortFromConfigObject sets httpPort variable to port mentioned in the config object
func setHTTPPortFromConfigObject() {
	httpPort = envs.ServerEnv.Port
}

// setEndpointFromConfigObject sets endpoint IP from config object
func setEndpointFromConfigObject() {
	if envs.ServerEnv.IsTLSEnabled {
		prefix = "https"
	} else {
		prefix = "http"
	}
	endpoint = fmt.Sprintf("%s:%s", envs.ServerEnv.Uri, httpPort)
}

// setTimeoutsFromConfigObject sets writeTimeout and readTimeout from config object
func setTimeoutsFromConfigObject() {
	writeTimeout = time.Duration(envs.ServerEnv.WriteTimeout) * time.Second
	readTimeout = time.Duration(envs.ServerEnv.ReadTimeout) * time.Second
}
