package bookcar

import (
	"time"

	"github.com/cv65kr/saga-temporal/sdk"
	"go.temporal.io/sdk/workflow"
)

type AwaitSignal struct {
	ConfirmationSignal bool
	PollStatusResult   string
}

func BookCarWorkflow(ctx workflow.Context, booking sdk.Booking) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("BookCarWorkflow workflow started", "booking id", booking.Id)

	signals := &AwaitSignal{}
	workflow.Go(ctx, signals.PollStatus)
	workflow.Go(ctx, signals.SignalListener)

	// Wait for Signal
	err := workflow.Await(ctx, func() bool {
		return signals.ConfirmationSignal
	})

	// Cancellation
	if err != nil {
		return err
	}

	logger.Info("BookCarWorkflow workflow completed.")

	return nil
}

func (a *AwaitSignal) PollStatus(ctx workflow.Context) {
	log := workflow.GetLogger(ctx)
	var result string

	// Infiity loop to have always fresh data about available slots
	for {
		err := workflow.ExecuteActivity(ctx, PollStatusActivity).Get(ctx, &result)
		if err != nil {
			log.Error("Activity failed.", "Error", err)
			continue
		}

		a.PollStatusResult = result

		err = workflow.NewTimer(ctx, time.Minute*5).Get(ctx, nil)
		if err != nil {
			break
		}
	}
}

func (a *AwaitSignal) SignalListener(ctx workflow.Context) {
	log := workflow.GetLogger(ctx)
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(workflow.GetSignalChannel(ctx, "ConfirmationSignal"), func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, nil)
			a.ConfirmationSignal = true
			log.Info("Confirmation signal received")
		})
		selector.Select(ctx)
	}
}

func BookCarCompensationWorkflow(ctx workflow.Context, booking sdk.Booking) error {
	return nil
}
