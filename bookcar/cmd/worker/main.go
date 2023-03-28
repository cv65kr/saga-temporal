package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/cv65kr/saga-temporal/bookcar"
	"github.com/cv65kr/saga-temporal/sdk"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, sdk.BookCarTaskQueue, worker.Options{})

	w.RegisterWorkflow(bookcar.BookCarWorkflow)
	w.RegisterWorkflow(bookcar.BookCarCompensationWorkflow)
	w.RegisterActivity(bookcar.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
