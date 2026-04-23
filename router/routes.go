package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marceloxhenrique/gopportunities/config"
	"github.com/marceloxhenrique/gopportunities/docs"
	"github.com/marceloxhenrique/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// handler.InitializeHandler()
	db := config.GetSQLite()
	h := handler.NewHandler(db)
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", h.ShowOpeningHandler)
		v1.POST("/opening", h.CreateOpeningHandler)
		v1.PUT("/opening", h.UpdateOpeningHandler)
		v1.DELETE("/opening", h.DeleteOpeningHandler)
		v1.GET("/openings", h.ListOpeningsHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
