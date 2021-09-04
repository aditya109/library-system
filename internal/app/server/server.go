package server

import (
	"fmt"
	"library-server/internal/app/server/router"
	"library-server/configs"
	"library-server/pkg/logger"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type ServerEnv struct {
	Address      string
	WriteTimeout int
	ReadTimeout  int
}

var standardLogger *logger.StandardLogger
var config configs.Config
var serverEnv *ServerEnv
var errorMessage string
var err error

func runSideCar() (string, error) {
	// initializing standard logger
	standardLogger = logger.NewLogger()
	// loading configuration
	filePath, err := filepath.Abs("./configs/config.json")

	if err != nil {
		return "error while getting config file path", err
	}
	filePath = strings.Split(filePath, " <nil>")[0]
	config, errorMessage, err = configs.LoadConfiguration(filePath, standardLogger)
	if err != nil {
		return errorMessage, err
	}

	serverEnv = new(ServerEnv)
	if strings.ToUpper(config.EnvironmentType) == "PROD" {

	} else if strings.ToUpper(config.EnvironmentType) == "DEV" {
		serverEnv.Address = fmt.Sprintf("%s:%s", config.DevEnvs.ServerUri, config.DevEnvs.ServerPort)
		serverEnv.WriteTimeout = config.DevEnvs.WriteTimeout
		serverEnv.ReadTimeout = config.DevEnvs.ReadTimeout
	}
	return "", nil
}

func StartServer() {
	// running side car
	errorMessage, err = runSideCar()
	if err != nil {
		log.Fatal("could not get config")
	}

	// initializing mux router
	router := router.GetRouter(config.APIPrefix, config)
	// initializing http server
	srv := &http.Server{
		Handler:      router,
		Addr:         serverEnv.Address,
		WriteTimeout: time.Duration(serverEnv.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(serverEnv.ReadTimeout) * time.Second,
	}
	standardLogger.ServerEvent(fmt.Sprintf("server starting at %s", serverEnv.Address))
	// starting server
	err = srv.ListenAndServe()
	if err != nil {
		standardLogger.Issue("couldn't start server")
	}

}
