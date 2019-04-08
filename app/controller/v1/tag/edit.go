package tag

import (
	"github.com/gin-gonic/gin"
	"leest/app/service/v1/tag"
)

func Edit(c *gin.Context) {
	tag.Edit(c)
}