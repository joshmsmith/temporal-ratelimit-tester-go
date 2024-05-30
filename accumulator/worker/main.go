package main

import (
	"log"

	accumulator "webapp/accumulator"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	u "webapp/utils"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	// Load the Temporal Cloud from env
	clientOptions, err := u.LoadClientOptions(u.NoSDKMetrics)
	if err != nil {
		log.Fatalf("Failed to load Temporal Cloud environment: %v", err)
	}
	log.Println("connecting to temporal server..")
	c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, "accumulate_greetings", worker.Options{})

	w.RegisterWorkflow(accumulator.AccumulateSignalsWorkflow)
	w.RegisterActivity(accumulator.ComposeGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
