package main

import (
	"context"
	"log"
	"math/rand"

	"go.temporal.io/sdk/client"

	u "webapp/utils"
)

var WorkflowIDPrefix = "accumulate"

var TaskQueue = "accumulate_greetings";

func main() {
	for i := 0; i<100000; i++ {
			query()
	}
}

func query() {
	var workflowID, queryType, bucket string


	// The client is a heavyweight object that should be created once per process.
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

    
    buckets := []string{"red", "blue", "green", "yellow", "orange", "brown", "white", "purple", "avocado", "black", "cherry", "skyblue"}
	bucketIndex := rand.Intn(len(buckets))
	bucket = buckets[bucketIndex]
	workflowID = WorkflowIDPrefix + "-" + bucket;
	queryType = "status"
	resp, err := c.QueryWorkflow(context.Background(), workflowID, "", queryType)
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}
	var result interface{}
	if err := resp.Get(&result); err != nil {
		log.Fatalln("Unable to decode query result", err)
	}
	log.Println("Received query result", "Result", result)

}