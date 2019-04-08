package article 

func ExistById(id int) bool {
	var article Article;

	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	
	return false
}