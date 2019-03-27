package tag 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"leest/infra/e"
	"leest/domain/entity"
)

func Add(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		code = e.SUCCESS
		entity.AddTag(name, state, createBy)
	}

	c.JSON(http.StatusOK, gin.H {
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
