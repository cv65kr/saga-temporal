package main

import (
	"log"
	"math/rand"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/cv65kr/saga-temporal/bookflight"
	"github.com/cv65kr/saga-temporal/sdk"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, sdk.BookFlightTaskQueue, worker.Options{})

	w.RegisterWorkflow(bookflight.BookFlightWorkflow)
	w.RegisterActivity(bookflight.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
