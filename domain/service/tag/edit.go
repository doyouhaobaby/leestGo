package tag 

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"leest/infra/e"
	"leest/domain/entity/tag"
)

func Edit(c *gin.Context) map[string]interface{} {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许 0 或者 1")
	}

	valid.Required(id, "id").Message("ID 不能为空")
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100 个字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为 100 个字符")
	
	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if tag.ExistById(id) {
			code = e.SUCCESS

			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy

			if name != "" {
				data["name"] = name
			}

			if state != -1 {
				data["state"] = state
			}

			tag.Update(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	return map[string]interface{} {
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	}
}
