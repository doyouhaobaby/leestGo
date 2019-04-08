package tag

import (
	"net/http"
	"leest/domain/service/tag"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, tag.Destroy(c))
}
