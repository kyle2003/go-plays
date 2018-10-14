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

func FetchUnReapedCategoryList() []models.Category {
	db := conf.GlobalDb.Get()
	var categories []models.Category

	db.Model(&models.Category{}).Where("f_reap_status=?", 2).Scan(&categories)
	return categories
}

func FetchSubjectList() []models.Subject {
	db := conf.GlobalDb.Get()
	var subjs []models.Subject

	db.Model(&models.Subject{}).Scan(&subjs)
	return subjs
}

func FetchReapedSubjectList() []models.Subject {
	db := conf.GlobalDb.Get()
	var subjs []models.Subject

	db.Model(&models.Subject{}).Where("f_reap_status=? and f_download_status=?", 1, 2).Scan(&subjs)
	return subjs
}

func FetchSubjectsByCategoryID(cID uint64) []models.Subject {
	return nil
}

func FetchThumbImageBySubjectID(sID uint64) uint64 {
	db := conf.GlobalDb.Get()
	var img models.Image
	db.Where("f_category_id=?", sID).First(&img)
	return img.ID
}

func GetCategoryTitleByID(cID uint64) string {
	db := conf.GlobalDb.Get()
	var c models.Category
	db.Where("f_id=?", cID).First(&c)
	return c.Title
}

func GetImagesBySubjectID(sID uint64) []models.Image {
	db := conf.GlobalDb.Get()
	var images []models.Image
	db.Model(&models.Image{}).Where("f_subject_id=?", sID).Scan(&images)
	return images
}

func GetNotDownloadedImagesBySubjectID(sID uint64) []models.Image {
	db := conf.GlobalDb.Get()
	var images []models.Image
	db.Model(&models.Image{}).Where("f_subject_id=? and f_download_status=?", sID, 2).Scan(&images)
	return images
}
