package saga

import (
	"github.com/cv65kr/saga-temporal/sdk"
	"go.temporal.io/sdk/workflow"
)

func BookCarWorkflowSuccessMock(ctx workflow.Context, booking sdk.Booking) error {
	return nil
}

func BookHotelWorkflowSuccessMock(ctx workflow.Context, booking sdk.Booking) error {
	return nil
}

func BookFlightWorkflowSuccessMock(ctx workflow.Context, booking sdk.Booking) error {
	return nil
}
