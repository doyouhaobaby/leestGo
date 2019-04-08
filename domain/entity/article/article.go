package article 

import (
    "time"
    "github.com/jinzhu/gorm"
    "leest/domain/entity"
    "leest/domain/entity/tag"
)

type Article struct {
    entity.Entity

    TagID int `json:"tag_id" gorm:"index"`
    Tag   tag.Tag `json:"tag"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

var db = entity.Db

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}