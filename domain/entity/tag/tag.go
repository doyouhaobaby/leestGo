package tag

import (
    "time"
    "github.com/jinzhu/gorm"
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

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}