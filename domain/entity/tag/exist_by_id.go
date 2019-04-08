package tag

func ExistById(id int) bool {
	var tag Tag;

	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}
	
	return false
}