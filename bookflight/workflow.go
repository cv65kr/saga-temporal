package bookflight

import (
	"time"

	"github.com/cv65kr/saga-temporal/sdk"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func BookFlightWorkflow(ctx workflow.Context, booking sdk.Booking) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 2,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("BookFlightWorkflow workflow started", "booking id", booking.Id)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, booking).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("BookFlightWorkflow workflow completed.", "result", result)

	return nil
}
