// +build integration

package integration_test

import (
	"context"
	"sync"
	"time"

	"github.com/chef/automate/lib/workflow"
)

// TestEnqueueDuplicateWorkflowInstance tries the enqueue the same workflow
// instance twice. The first one must be accepted, the second one must fail
func (suite *WorkflowTestSuite) TestEnqueueDuplicateWorkflowInstance() {
	taskName := randName("duplicate_test")
	workflowName := randName("duplicate_test")
	instanceName := randName("instance")

	wgTask := sync.WaitGroup{}
	wgTask.Add(1)

	// wgBarrier is used to block the completion of the task.
	// This allows us to make sure there is a running workflow
	// instance
	wgBarrier := sync.WaitGroup{}
	wgBarrier.Add(1)

	m := suite.newManager(
		WithTaskExecutorF(
			taskName,
			func(context.Context, workflow.Task) (interface{}, error) {
				wgBarrier.Wait()
				return nil, nil
			}),
		WithWorkflowExecutor(
			workflowName,
			&workflowExecutorWrapper{
				onStart: func(w workflow.WorkflowInstance, ev workflow.StartEvent) workflow.Decision {
					err := w.EnqueueTask(taskName, nil)
					suite.Require().NoError(err, "failed to enqueue task")
					return w.Continue(nil)
				},
				onTaskComplete: func(w workflow.WorkflowInstance, ev workflow.TaskCompleteEvent) workflow.Decision {
					wgTask.Done()
					return w.Complete()
				},
			},
		),
	)
	defer m.Stop()
	err := m.EnqueueWorkflow(context.Background(), workflowName, instanceName, nil)
	suite.Require().NoError(err, "Failed to enqueue workflow")
	err = m.EnqueueWorkflow(context.Background(), workflowName, instanceName, "dummy")
	suite.Assert().Equal(workflow.ErrWorkflowInstanceExists, err)
	wgBarrier.Done()
	wgTask.Wait()
	time.Sleep(20 * time.Millisecond)
	err = m.Stop()
	suite.NoError(err)
}
