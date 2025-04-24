package main

import (
	"log"
	"os"
	"time"

	"github.com/wiredlush/easy-gate/internal/config"
	"github.com/wiredlush/easy-gate/internal/engine"
	"github.com/wiredlush/easy-gate/internal/routine"
)

func main() {
	cfgFilePath, err := config.GetConfigPath(os.Args)
	if err != nil {
		log.Fatal("No configuration file provided")
	}

	log.Println("Loading configuration file:", cfgFilePath)
	cfgRoutine, err := routine.NewRoutine(cfgFilePath, 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	go cfgRoutine.Start()

	engine.NewEngine(cfgRoutine).Serve()
}
