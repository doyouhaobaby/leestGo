package article 

func Gets(pageNum int, pageSize int, maps interface {}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}
