package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/why-xn/go-temporal-skeleton/pkg/server/router"
)

func Start() {
	r := gin.Default()

	// Setting API Base Path for HTTP APIs
	httpRouter := r.Group("/")

	// Setting up all Http Routes
	router.AddApiRoutes(httpRouter)

	// Setup Swagger Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("[INFO]", "Starting Server...")
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Println("[ERROR]", err)
	}
}
