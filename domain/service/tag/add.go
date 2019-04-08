package tag 

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"leest/infra/e"
	"leest/domain/entity/tag"
)

func Add(c *gin.Context) map[string]interface{} {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100 个字符")
	valid.Required(createBy, "create_by").Message("创建人不能为空")
	valid.MaxSize(createBy, 100, "create_by").Message("创建人最长为 100 个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许 0 或者 1")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if tag.ExistByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			code = e.SUCCESS
			tag.Add(name, state, createBy)
		}
	}

	return map[string]interface{} {
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	}
}
