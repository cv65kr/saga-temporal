package bookcar

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func BookCarWorkflow(ctx workflow.Context, name string) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("BookCarWorkflow workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	logger.Info("BookCarWorkflow workflow completed.", "result", result)

	return nil
}

func BookCarCompensationWorkflow(ctx workflow.Context, name string) error {
	return nil
}
