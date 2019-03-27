package tag

import (
	"leest/app/service/v1/tag"
	"github.com/gin-gonic/gin"
)

func Gets(c *gin.Context) {
	tag.Gets(c)
}
