package service

import (
	ctx "context"
)

type ApiServiceInterface interface {
	CheckPermission(c ctx.Context) error
	Validate(c ctx.Context) error
	Process(c ctx.Context) error
	PostProcess(c ctx.Context) error
}
