package routers 

import (
	"github.com/gin-gonic/gin"
	"leest/infra/setting"
	v1tag "leest/app/controller/v1/tag"
	v1article "leest/app/controller/v1/article"
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

		apiv1.GET("/articles", v1article.Gets)
		apiv1.GET("/articles/:id", v1article.Show)
		apiv1.POST("/articles", v1article.Store)
		apiv1.PUT("/articles/:id", v1article.Update)
		apiv1.DELETE("/articles/:id", v1article.Destroy)
	}

	return router
}