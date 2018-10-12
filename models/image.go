package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Image struct {
	// generic attributes of the image object
	PandoraObj
	// categoryID indicates which category image belongs to
	CategoryID uint64 `gorm:"column:f_category_id;index;" json:"category_id"`
	// subjectID indicates where the image belongs to
	SubjectID uint64 `gorm:"column:f_subject_id;index;" json:"subject_id"`
	// base64 string, img content
	Base64 string `gorm:"column:f_base64;type:text;" json:"base64"`
}

// Create db Record
func (i *Image) Create(db *gorm.DB) error {
	// default attributes
	i.Created = time.Now().Unix()
	i.Updated = time.Now().Unix()

	db.AutoMigrate(i)
	return db.Create(i).Error
}
