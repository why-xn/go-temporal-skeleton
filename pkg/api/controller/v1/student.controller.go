package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/why-xn/go-temporal-skeleton/pkg/api"
	servicev1 "github.com/why-xn/go-temporal-skeleton/pkg/api/service/student/v1"
	workflowv1 "github.com/why-xn/go-temporal-skeleton/pkg/api/workflow/student/v1"
	"github.com/why-xn/go-temporal-skeleton/pkg/core/log"
	"github.com/why-xn/go-temporal-skeleton/pkg/core/temporal"
	"go.temporal.io/sdk/client"
	"time"
)

type StudentControllerInterface interface {
	GetStudentRules(c *gin.Context)
}

type studentController struct{}

var ac studentController

func StudentController() *studentController {
	return &ac
}

var studentWorkflowOptions = client.StartWorkflowOptions{
	ID:                       uuid.NewString(),
	TaskQueue:                temporal.WorkerQueue,
	WorkflowExecutionTimeout: time.Second * 10,
}

// GetStudentList godoc
// @Summary      Get Student List
// @Description  Get Student List
// @Tags         student
// @Param request body v1.GetGetStudentsInputParams true "input"
// @Accept       json
// @Produce      json
// @Success      200  {object}  api.ResponseDTO
// @Failure      400  {object}  api.ResponseDTO
// @Security	 ApiKeyAuth
// @Router       /v1/students [get]
func (ctrl *studentController) GetStudentList(ctx *gin.Context) {
	var result api.ResponseDTO

	input := new(servicev1.GetStudentListInputParams)
	err := ctx.BindQuery(input)
	if err != nil {
		log.Logger.Errorw("[ERROR]", "err", err)
		return
	}

	we, err := temporal.Client().ExecuteWorkflow(ctx, studentWorkflowOptions, workflowv1.StudentWorkflow().GetStudentList, *input)
	if err != nil {
		log.Logger.Errorw("Unable to start the GetStudentRules workflow", "err", err, "wid", we.GetID())
		api.SendErrorResponse(ctx, err.Error())
		return
	}

	log.Logger.Debugw("GetStudentRules workflow started.", "wid", we.GetID(), "rid", we.GetRunID())

	err = we.Get(ctx, &result)
	if err != nil {
		log.Logger.Errorw("Unable to get result for GetStudentRules", "err", err, "wid", we.GetID(), "rid", we.GetRunID())
	}

	api.SendResponse(ctx, result)
}

// GetStudent godoc
// @Summary      Get Student List
// @Description  Get Student List
// @Tags         student
// @Param request body v1.GetStudentInputParams true "input"
// @Accept       json
// @Produce      json
// @Success      200  {object}  api.ResponseDTO
// @Failure      400  {object}  api.ResponseDTO
// @Security	 ApiKeyAuth
// @Router       /v1/students/:id [get]
func (ctrl *studentController) GetStudent(ctx *gin.Context) {
	var result api.ResponseDTO

	input := new(servicev1.GetStudentInputParams)
	input.Id = ctx.Param("id")

	we, err := temporal.Client().ExecuteWorkflow(ctx, studentWorkflowOptions, workflowv1.StudentWorkflow().GetStudent, *input)
	if err != nil {
		log.Logger.Errorw("Unable to start the GetStudentRules workflow", "err", err, "wid", we.GetID())
		api.SendErrorResponse(ctx, err.Error())
		return
	}

	log.Logger.Debugw("GetStudentRules workflow started.", "wid", we.GetID(), "rid", we.GetRunID())

	err = we.Get(ctx, &result)
	if err != nil {
		log.Logger.Errorw("Unable to get result for GetStudentRules", "err", err, "wid", we.GetID(), "rid", we.GetRunID())
	}

	api.SendResponse(ctx, result)
}
