package main

import (
	"flag"
	"github.com/why-xn/go-temporal-skeleton/pkg/core/log"
	"github.com/why-xn/go-temporal-skeleton/pkg/core/temporal"
	"github.com/why-xn/go-temporal-skeleton/pkg/server"
)

func main() {
	// Initialize Zap logger
	log.InitializeLogger()

	// Initializing Temporal Client
	err := temporal.InitClient()
	if err != nil {
		log.Logger.Fatal(err)
	} else {
		log.Logger.Info("Temporal Client Connected")
	}
	defer temporal.CloseClient()

	var serviceType string
	flag.StringVar(&serviceType, "service-type", "full", "define the service type based on the flag")
	flag.Parse()

	startService(serviceType)

}

func startService(serviceType string) {
	log.Logger.Info("Service Type: ", serviceType)

	if serviceType == "" || serviceType == "full" {
		// Start All Temporal Workers
		go temporal.StartWorker(temporal.Client())

		// Start Gin Server
		server.Start()

	} else if serviceType == "worker" {
		temporal.StartWorker(temporal.Client())

	} else if serviceType == "api-gateway" {
		server.Start()
	}
}
