package tag

import (
	"net/http"
	"leest/domain/service/tag"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, tag.Edit(c))
}
