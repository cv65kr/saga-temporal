package saga

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

func Test_SagaWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.RegisterWorkflowWithOptions(BookCarWorkflowSuccessMock, workflow.RegisterOptions{
		Name: "BookCarWorkflow",
	})

	env.RegisterWorkflowWithOptions(BookHotelWorkflowSuccessMock, workflow.RegisterOptions{
		Name: "BookHotelWorkflow",
	})

	env.RegisterWorkflowWithOptions(BookFlightWorkflowSuccessMock, workflow.RegisterOptions{
		Name: "BookFlightWorkflow",
	})

	env.ExecuteWorkflow(SagaWorkflow, "Temporal")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}
