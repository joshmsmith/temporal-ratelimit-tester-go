package main

import (
	"context"
	"fmt"
	"log"
	//"time"

	"math/rand"

	accumulator "webapp/accumulator"
	"go.temporal.io/sdk/client"
	
	u "webapp/utils"
)


var WorkflowIDPrefix = "accumulate"

var TaskQueue = "accumulate_greetings";
func main() {
	// The client is a heavyweight object that should be created once per process.
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


	// setup to send signals
    bucket := "blue";
    workflowId := WorkflowIDPrefix + "-" + bucket;
    buckets := []string{"red", "blue", "green", "yellow", "orange", "brown", "white", "purple", "avocado", "black", "cherry", "skyblue"}
    names := []string{"Genghis Khan", "Missy", "Bill", "Ted", "Rufus", "Abe"}
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowId,
		TaskQueue: TaskQueue,
	}

	// 2000 signals/20ms delay is about 15 APS
	// 10 ms is about 16 APS
	// 20000 signals/ 5 ms/ 1 starter is about 22 APS, 65 RPS
	max_signals := 200000

    for i := 0; i < max_signals; i++ {
		bucketIndex := rand.Intn(len(buckets))
		bucket = buckets[bucketIndex]
		nameIndex := rand.Intn(len(names))
		
		greeting := accumulator.AccumulateGreeting{
			GreetingText: names[nameIndex],
			Bucket: bucket,
			GreetingKey: "key-" + fmt.Sprint(i),
		}
		//time.Sleep(5 * time.Millisecond)

		workflowId = WorkflowIDPrefix + "-" + bucket
		workflowOptions = client.StartWorkflowOptions{
			ID:        workflowId,
			TaskQueue: TaskQueue,
		}
		we, err := c.SignalWithStartWorkflow(context.Background(), workflowId, "greeting", greeting, workflowOptions, accumulator.AccumulateSignalsWorkflow, bucket)
		if err != nil {
			log.Fatalln("Unable to signal with start workflow", err)
		}
		var progress float32
		progress = 100.0 * float32(i)/float32(max_signals)
		log.Println("[" + fmt.Sprintf("%.2f%%",progress) + "] Signaled/Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID(), "signal:", greeting.GreetingText)

    }	
	
	return


}
