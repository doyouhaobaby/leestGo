package article

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"leest/infra/e"
	"leest/domain/entity/article"
	"leest/infra/util"
	"leest/infra/setting"
)

func Lists(c *gin.Context) map[string]interface{} {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态值只允许 0 或者 1")
	}

	var tagId int = -1

	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签 ID 必须大于 0")
	}

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		code = e.SUCCESS
		
		data["lists"] = article.Lists(util.GetPage(c), setting.PageSize, maps)
		data["total"] = article.Count(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	return map[string]interface{}{
		"code": code,
		"msg": e.GetMsg(code),	
		"data": data,
	}
}
