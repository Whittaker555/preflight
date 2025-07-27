package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/preflight/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/health", handlers.HealthCheck)
	api.POST("/plan/analyse", handlers.AnalysePlan)
	api.POST("/plan/upload", handlers.UploadPlan)
	api.POST("/plan/result/upload", handlers.UploadResult)
}
