package article

import (
	"net/http"
	"leest/infra/repository/article"
	"github.com/gin-gonic/gin"
)

func Gets(c *gin.Context) {
	c.JSON(http.StatusOK, article.Lists(c))
}
