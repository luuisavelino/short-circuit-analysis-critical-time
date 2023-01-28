package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-critical-time/controllers"
	"github.com/luuisavelino/short-circuit-analysis-critical-time/middleware"
)

func HandleRequests() {
	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/actuator/health"),
		gin.Recovery(),
		middleware.Logger(),
	)

	actuator := router.Group("/actuator")
	{
		actuator.GET("/health", controllers.HealthGET)
	}

	systemData := router.Group("/api/v2/files/:fileId/short-circuit/:line/point/:point")
	{
		systemData.GET("/fault", controllers.AllData)
		systemData.GET("/fault/:gerador", controllers.GeradorData)
	}

	router.Run(":8082")
}
