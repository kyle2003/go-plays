package models

import (
	"errors"
	"pandora/constants"
	"pandora/utils"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Subject struct {
	// generic attributes
	PandoraObj
	// categoryID where subject belongs to
	CategoryID uint64 `gorm:"column:f_category_id;index;" json:""`
	// images collection num
	ImagesNum int `gorm:"column:f_images_num;default:0;" json:""`
	// the thumb imageid
	ThumbImageID uint64 `gorm:"column:f_thumb_image_id;default:0;" json:""`
	// images object collection
	Images []Image `gorm:"-" json:"-"`
}

// Create db
func (s *Subject) Create(db *gorm.DB) error {
	// default attributes
	s.Created = time.Now().Unix()
	s.Updated = time.Now().Unix()

	db.AutoMigrate(s)
	return db.Create(s).Error
}

func (s *Subject) ReapImages(db *gorm.DB) error {
	html := s.GetHTML(0)
	reg, _ := regexp.Compile(`img src="//(.*.jpg)"`)
	strs := reg.FindAllString(html, -1)
	urlsStr := ""
	for _, str := range strs {
		urlsStr += str
	}

	reg, _ = regexp.Compile(`img|alt`)
	regImg, _ := regexp.Compile(`.*//(.*.jpg).*`)
	repl := "${1}"

	// 截取到的字符串
	for _, str := range reg.Split(urlsStr, -1) {
		if match, _ := regexp.MatchString(".*jpg", str); match {
			var img Image
			str = strings.Replace(str, " ", "", -1)
			url := "http://" + regImg.ReplaceAllString(str, repl)

			if m, _ := regexp.MatchString(".jpg$", url); m {
				img.Name = utils.Basename(str)
				img.URL = url
				img.Title = s.Title
				img.SubjectID = s.ID
				img.CategoryID = s.CategoryID
				img.ReapStatus = constants.REAP_STATUS__DONE
				s.Images = append(s.Images, img)
				img.Create(db)
				s.ImagesNum++
			}
		}
	}

	// If 0 images reaped, then disable the subject
	if s.ImagesNum == 0 {
		s.Enabled = constants.BOOL__FALSE
		return errors.New("Reaped 0 image links for " + s.Title)
	}
	return nil
}
