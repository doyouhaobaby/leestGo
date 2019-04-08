package tag

func Destroy(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}
