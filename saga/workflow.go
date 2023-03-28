package saga

import (
	"github.com/cv65kr/saga-temporal/sdk"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
)

func SagaWorkflow(ctx workflow.Context, name string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("SAGA workflow started", "name", name)

	// Execute BOOK CAR
	cwo := workflow.ChildWorkflowOptions{
		TaskQueue: sdk.BookCarTaskQueue,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	err := workflow.ExecuteChildWorkflow(ctx, "BookCarWorkflow", name).Get(ctx, nil)
	if err != nil {
		logger.Error("BookCarWorkflow failed.", "Error", err)
		return err
	}

	// Compensation
	defer func() {
		if err != nil {
			cwo := workflow.ChildWorkflowOptions{
				TaskQueue: sdk.BookCarTaskQueue,
			}
			ctx = workflow.WithChildOptions(ctx, cwo)
			errCompensation := workflow.ExecuteChildWorkflow(ctx, "BookCarCompensationWorkflow", name).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	// Execute BOOK HOTEL
	cwo = workflow.ChildWorkflowOptions{
		TaskQueue: sdk.BookHotelTaskQueue,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	err = workflow.ExecuteChildWorkflow(ctx, "BookHotelWorkflow", name).Get(ctx, nil)
	if err != nil {
		logger.Error("BookHotelWorkflow failed.", "Error", err)
		return err
	}

	// Compensation
	defer func() {
		if err != nil {
			cwo = workflow.ChildWorkflowOptions{
				TaskQueue: sdk.BookHotelTaskQueue,
			}
			ctx = workflow.WithChildOptions(ctx, cwo)
			errCompensation := workflow.ExecuteChildWorkflow(ctx, "BookHotelCompensationWorkflow", name).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	// Execute BOOK FLIGHT
	cwo = workflow.ChildWorkflowOptions{
		TaskQueue: sdk.BookFlightTaskQueue,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	err = workflow.ExecuteChildWorkflow(ctx, "BookFlightWorkflow", name).Get(ctx, nil)
	if err != nil {
		logger.Error("BookFlightWorkflow failed.", "Error", err)
		return err
	}

	return err
}
