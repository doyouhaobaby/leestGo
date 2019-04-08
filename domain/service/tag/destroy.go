package tag 

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"leest/infra/e"
	"leest/domain/entity/tag"
)

func Destroy(c *gin.Context) map[string]interface{} {
	id := com.StrTo(c.Param("id")).MustInt()
	
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID 不能为空")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if tag.ExistById(id) {
			code = e.SUCCESS

			tag.Destroy(id)
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
