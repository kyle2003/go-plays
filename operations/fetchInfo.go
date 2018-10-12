package operations

import (
	"pandora/conf"
	"pandora/models"
)

func FetchCategoryList() []models.Category {
	db := conf.GlobalDb.Get()
	var categories []models.Category

	db.Model(&models.Category{}).Scan(&categories)
	return categories
}

func FetchSubjectsByCategoryID(cID uint64) []models.Subject {
	return nil
}
