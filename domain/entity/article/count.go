package article 

func Count(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}
