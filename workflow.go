package longworkflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func LongHelloWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

 	var result string

  for range(12) {
    workflow.Sleep(time.Second * 10)
  	err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &result)
  	if err != nil {
   		return "", err
  	}
  }

	return result, nil
}

