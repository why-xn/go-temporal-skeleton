package v1

import (
	ctx "context"
	"errors"
	"github.com/google/uuid"
	"github.com/why-xn/go-temporal-skeleton/pkg/api"
	. "github.com/why-xn/go-temporal-skeleton/pkg/api/service"
	"github.com/why-xn/go-temporal-skeleton/pkg/types"
)

type StudentServiceInterface interface {
	GetStudentList(c ctx.Context, params GetStudentListInputParams) (interface{}, error)
	GetStudent(c ctx.Context, params GetStudentInputParams) (interface{}, error)
}

type studentService struct{}

var as studentService

func StudentService() *studentService {
	return &as
}

// ---- Get Student List ---- //

type GetStudentListInputParams struct {
	output []types.Student
	BaseInternalParams
}

func (svc *studentService) GetStudentList(c ctx.Context, p GetStudentListInputParams) (interface{}, error) {
	err := p.CheckPermission(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	err = p.Validate(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	err = p.Process(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	_ = p.PostProcess(c)

	return api.SuccessResponse(p.output)
}

func (p *GetStudentListInputParams) CheckPermission(c ctx.Context) error {
	return nil
}

func (p *GetStudentListInputParams) Validate(c ctx.Context) error {
	return nil
}

func (p *GetStudentListInputParams) Process(c ctx.Context) error {
	p.output = []types.Student{
		{
			Id:        uuid.NewString(),
			Name:      "Shihab Hasan",
			StudentId: "123456",
		},
		{
			Id:        uuid.NewString(),
			Name:      "Nahid Hasan",
			StudentId: "654321",
		},
	}
	return nil
}

func (p *GetStudentListInputParams) PostProcess(c ctx.Context) error {
	return nil
}

// ---- Get Student ---- //

type GetStudentInputParams struct {
	Id string `json:"id"`

	output types.Student
	BaseInternalParams
}

func (svc *studentService) GetStudent(c ctx.Context, p GetStudentInputParams) (interface{}, error) {
	err := p.CheckPermission(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	err = p.Validate(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	err = p.Process(c)
	if err != nil {
		return api.ErrorResponse(err)
	}

	_ = p.PostProcess(c)

	return api.SuccessResponse(p.output)
}

func (p *GetStudentInputParams) CheckPermission(c ctx.Context) error {
	return nil
}

func (p *GetStudentInputParams) Validate(c ctx.Context) error {
	if p.Id == "" {
		return errors.New("id is missing")
	}
	return nil
}

func (p *GetStudentInputParams) Process(c ctx.Context) error {
	p.output = types.Student{
		Id:        uuid.NewString(),
		Name:      "Shihab Hasan",
		StudentId: "123456",
	}
	return nil
}

func (p *GetStudentInputParams) PostProcess(c ctx.Context) error {
	return nil
}
