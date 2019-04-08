package tag

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"leest/infra/e"
	"leest/domain/entity/tag"
	"leest/infra/util"
	"leest/infra/setting"
)

func Gets(c *gin.Context) map[string]interface{} {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = tag.Gets(util.GetPage(c), setting.PageSize, maps)
	data["total"] = tag.GetTotal(maps)

	return map[string]interface{}{
		"code": code,
		"msg": e.GetMsg(code),	
		"data": data,
	}
}
