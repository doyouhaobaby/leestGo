package tag

import (
	"leest/domain/entity"
)

type Tag struct {
    entity.Entity

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

var db = entity.Db