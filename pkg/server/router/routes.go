package router

import (
	"github.com/gin-gonic/gin"
	"github.com/why-xn/go-temporal-skeleton/pkg/api/controller/v1"
)

// @title           Swagger API
// @version         1.0

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func AddApiRoutes(httpRg *gin.RouterGroup) {
	//httpRg.Use(auth.TokenAuthMiddleware())
	httpRg.GET("api/v1/students", v1.StudentController().GetStudentList)
	httpRg.GET("api/v1/students/:id", v1.StudentController().GetStudent)
}
