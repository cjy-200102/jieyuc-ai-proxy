package routes

import (
	"github.com/gin-gonic/gin"
	"jieyuc-ai-proxy/handler"
	"net/http"
)

func InitHttpRoutes(r *gin.Engine) {
	group := r.Group("/jieyuc-ai-proxy")
	group.POST("/api/v1/exec", handler.DeepseekExec)
	group.GET("/api/v1/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
