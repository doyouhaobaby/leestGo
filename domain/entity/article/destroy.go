package article 

func Destroy(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})

	return true
}
