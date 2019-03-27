package tag

import (
	"net/http"
	"leest/infra/repository/tag"
	"github.com/gin-gonic/gin"
)

func Gets(c *gin.Context) {
	c.JSON(http.StatusOK, tag.Gets(c))
}
