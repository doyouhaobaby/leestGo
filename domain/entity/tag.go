package entity 

type Tag struct {
    Entity

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag {
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})

	return true
}
