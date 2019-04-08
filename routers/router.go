package routers 

import (
	"github.com/gin-gonic/gin"
	"leest/infra/setting"
	v1tag "leest/app/controller/v1/tag"
)

func InitRouter () *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())	

	gin.SetMode(setting.RunMode)

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/tags", v1tag.Gets)
		apiv1.POST("/tags", v1tag.Add)
		apiv1.PUT("/tags/:id", v1tag.Edit)
		apiv1.DELETE("/tags/:id", v1tag.Delete)
	}

	return router
}