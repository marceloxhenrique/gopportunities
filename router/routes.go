package router

import (
	"gopportunites/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
		v1.GET("/openings", handler.ListOpeningsHanlder)
	}
}
