package main

import (
	"context"
	"log"

	"github.com/cv65kr/saga-temporal/saga"
	"github.com/cv65kr/saga-temporal/sdk"
	"github.com/google/uuid"
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

	booking := sdk.Booking{
		Id: uuid.New().String(),
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, saga.SagaWorkflow, booking)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

}
