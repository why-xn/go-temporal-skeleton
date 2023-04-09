package v1

import (
	"github.com/why-xn/go-temporal-skeleton/pkg/api"
	servicev1 "github.com/why-xn/go-temporal-skeleton/pkg/api/service/student/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

// RetryPolicy specifies how to automatically handle retries if an Activity fails.
var retryPolicy = &temporal.RetryPolicy{
	InitialInterval:    2 * time.Second,
	BackoffCoefficient: 2.0,
	MaximumInterval:    100 * time.Second,
	MaximumAttempts:    3,
}

var options = workflow.ActivityOptions{
	StartToCloseTimeout: time.Second * 5,
	RetryPolicy:         retryPolicy,
}

type StudentWorkflowInterface interface {
	GetStudentRules(ctx workflow.Context, input servicev1.GetStudentListInputParams) (api.ResponseDTO, error)
}

type studentWorkflow struct{}

var tw studentWorkflow

func StudentWorkflow() *studentWorkflow {
	return &tw
}

func (t studentWorkflow) GetStudentList(ctx workflow.Context, input servicev1.GetStudentListInputParams) (api.ResponseDTO, error) {
	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)
	var res api.ResponseDTO

	err := workflow.ExecuteActivity(ctx, servicev1.StudentService().GetStudentList, input).Get(ctx, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t studentWorkflow) GetStudent(ctx workflow.Context, input servicev1.GetStudentInputParams) (api.ResponseDTO, error) {
	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)
	var res api.ResponseDTO

	err := workflow.ExecuteActivity(ctx, servicev1.StudentService().GetStudent, input).Get(ctx, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
