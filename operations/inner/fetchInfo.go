package inner

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

	db.Model(&models.Category{}).Where("f_reap_status=? and f_enabled=?", 2, 1).Scan(&categories)
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

	db.Model(&models.Subject{}).Where("f_reap_status=? and f_download_status=? and f_enabled=?", 1, 2, 1).Scan(&subjs)
	return subjs
}

func FetchSubjectsByCategoryID(cID uint64) []models.Subject {
	db := conf.GlobalDb.Get()
	var subjs []models.Subject

	db.Model(&models.Subject{}).Where("f_category_id=? and f_enabled=?", cID, 1).Scan(&subjs)
	return subjs
}

func GetThumbImageBySubjectID(sID uint64) uint64 {
	db := conf.GlobalDb.Get()
	var img models.Image
	db.Where("f_subject_id=? and f_enabled=?", sID, 1).Last(&img)
	return img.ID
}

func GetCategoryByID(cID uint64) models.Category {
	db := conf.GlobalDb.Get()
	var c models.Category
	db.Where("f_id=?", cID).First(&c)
	return c
}

func GetCategoryTitleByID(cID uint64) string {
	db := conf.GlobalDb.Get()
	var c models.Category
	db.Where("f_id=? and f_enabled=?", cID, 1).First(&c)
	return c.Title
}

func GetImagesBySubjectID(sID uint64) []models.Image {
	db := conf.GlobalDb.Get()
	var images []models.Image
	db.Model(&models.Image{}).Where("f_subject_id=? and f_enabled=?", sID, 1).Scan(&images)
	return images
}

func GetNotDownloadedImagesBySubjectID(sID uint64) []models.Image {
	db := conf.GlobalDb.Get()
	var images []models.Image
	db.Model(&models.Image{}).Where("f_subject_id=? and f_download_status=? and f_enabled=?", sID, 2, 1).Scan(&images)
	return images
}

func GetDownloadedImagesBySubjectID(sID uint64) []models.Image {
	db := conf.GlobalDb.Get()
	var images []models.Image
	db.Model(&models.Image{}).Where("f_subject_id=? and f_download_status=? and f_enabled=?", sID, 1, 1).Scan(&images)
	return images
}

func GetSubjectByID(sID uint64) models.Subject {
	db := conf.GlobalDb.Get()
	var s models.Subject
	s.ID = sID

	db.Where(&s).First(&s)
	return s
}
