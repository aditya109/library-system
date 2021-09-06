package server

import (
	"context"
	"fmt"
	"library-server/configs"
	"library-server/internal/app/server/router"
	"library-server/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
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
	initializeLogger()                 // initializing standard logger
	if err = getConfig(); err != nil { // getting configs
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

func initializeLogger() {
	standardLogger = logger.NewLogger()
}

func getConfig() error {
	filePath, err := filepath.Abs("./configs/config.json") // loading configuration
	if err != nil {
		return err
	}
	fmt.Println(filePath)
	config, errorMessage, err = configs.LoadConfiguration(strings.Split(filePath, " <nil>")[0], standardLogger)
	if err != nil {
		return err
	}
	return nil
}

func StartServer() {
	if errorMessage, err = runSideCar(); err != nil { // running side car
		log.Fatal(errorMessage)
	}
	srv := &http.Server{ // initializing http server
		Handler:      router.GetRouter(config.APIPrefix, config), // initializing mux router
		Addr:         serverEnv.Address,
		WriteTimeout: time.Duration(serverEnv.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(serverEnv.ReadTimeout) * time.Second,
	}
	go func() {
		standardLogger.ServerEvent(fmt.Sprintf("server starting at %s", serverEnv.Address))
		if err = srv.ListenAndServe(); err != nil { // starting server
			standardLogger.ServerEvent("couldn't start server")
		}
	}()
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan                                                          // Block until we receive our signal.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // Create a deadline to wait for.
	defer cancel()
	srv.Shutdown(ctx)
	standardLogger.ServerEvent("Shutting down")
	os.Exit(0)
}
