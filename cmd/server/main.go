package main

import (
	"github.com/aditya109/library-system/internal/server"
	logCfg "github.com/aditya109/library-system/pkg/logger"
)

func main() {
	logCfg.InitializeLogging() // initializing logger
	server.Start()
}
