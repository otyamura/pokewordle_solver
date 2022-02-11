package search

import "gorm.io/gorm"

func SearchCorrect(db *gorm.DB, q string) *gorm.DB {
	return db.Where("name LIKE ?", "%"+q+"%")
}

func SearchPartial(db *gorm.DB, p []string) *gorm.DB {
	tmp := db
	for _, v := range p {
		tmp = tmp.Where("name LIKE ?", "%"+v+"%")
	}
	return tmp
}

func SearchGeneration(db *gorm.DB, l string, g string) *gorm.DB {
	return db.Where("generation >= ? AND generation <= ?", l, g)
}
