package tag

import (
	"github.com/gin-gonic/gin"
	"leest/app/service/v1/tag"
)

func Delete(c *gin.Context) {
	tag.Delete(c)
}
