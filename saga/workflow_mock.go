package saga

import "go.temporal.io/sdk/workflow"

func BookCarWorkflowSuccessMock(ctx workflow.Context, name string) error {
	return nil
}

func BookHotelWorkflowSuccessMock(ctx workflow.Context, name string) error {
	return nil
}

func BookFlightWorkflowSuccessMock(ctx workflow.Context, name string) error {
	return nil
}
