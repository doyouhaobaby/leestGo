package tag

import (
	"github.com/gin-gonic/gin"
	"leest/app/service/v1/tag"
)

func Add(c *gin.Context) {
	tag.Add(c)
}
