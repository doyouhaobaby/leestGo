package article

import (
	"github.com/gin-gonic/gin"
	"leest/app/service/v1/article"
)

func Update(c *gin.Context) {
	article.Gets(c)
}
