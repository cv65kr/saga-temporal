package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/cv65kr/saga-temporal/bookhotel"
	"github.com/cv65kr/saga-temporal/sdk"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, sdk.BookHotelTaskQueue, worker.Options{})

	w.RegisterWorkflow(bookhotel.BookHotelWorkflow)
	w.RegisterWorkflow(bookhotel.BookHotelCompensationWorkflow)
	w.RegisterActivity(bookhotel.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
