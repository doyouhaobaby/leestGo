package tag

import (
	"leest/app/service/v1/tag"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	tag.Edit(c)
}