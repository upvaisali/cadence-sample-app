package main

import (
	"fmt"

	"github.com/upvaisali/cadence-sample-app/app/adapters/cadenceadapter"
	"github.com/upvaisali/cadence-sample-app/app/config"

	// "github.com/upvaisali/cadence-sample-app/app/worker/workflow"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

func startWorkers(h *cadenceadapter.CadenceAdapter, taskList string) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}

	cadenceWorker := worker.New(h.ServiceClient, h.Config.Domain, taskList, workerOptions)
	err := cadenceWorker.Start()
	if err != nil {
		h.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}

func main() {
	fmt.Println("Starting Worker..")
	var appConfig config.AppConfig
	appConfig.Setup()
	var cadenceClient cadenceadapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)

	//startWorkers(&cadenceClient, workflow.)
	// The workers are supposed to be long running process that should not exit.
	select {}
}
