package bookhotel

import (
	"time"

	"github.com/cv65kr/saga-temporal/sdk"
	"go.temporal.io/sdk/workflow"
)

func BookHotelWorkflow(ctx workflow.Context, booking sdk.Booking) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("BookHotelWorkflow workflow started", "booking id", booking.Id)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, booking).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("BookHotelWorkflow workflow completed.", "result", result)

	return nil
}

func BookHotelCompensationWorkflow(ctx workflow.Context, booking sdk.Booking) error {
	return nil
}
