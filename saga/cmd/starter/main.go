package main

import (
	"context"
	"log"

	"github.com/cv65kr/saga-temporal/saga"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "saga-tq",
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, saga.SagaWorkflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

}
