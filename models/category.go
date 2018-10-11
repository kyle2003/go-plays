package models

import (
	"pandora/constants"
	"pandora/utils"

	"github.com/jinzhu/gorm"

	"regexp"
	"time"
)

type Category struct {
	// object的属性
	PandoraObj
	// download limit
	Limit int `gorm:"column:f_limit;default:5;" json:"limit"`
	// subject nums
	SubjectsNum int `gorm:"column:f_subjects_num;" json:"subjects_num"`
	// subjects
	Subjects []Subject `gorm:"-"`
}

// Create db
func (c *Category) Create(db *gorm.DB) error {
	// default attributes
	c.ReapStatus = constants.REAP_STATUS__NOTDONE
	c.DownloadStatus = constants.DOWNLOAD_STATUS__NOTDONE
	c.Created = time.Now().Unix()
	c.Updated = time.Now().Unix()

	db.AutoMigrate(c)
	return db.Create(c).Error
}

// GetHtml content of the category page
func (c *Category) GetHtml() string {
	url := constants.BASE + c.Name
	return string(utils.GetHtml(url))
}

// GetPageLimit get the limit of page
func (c *Category) GetPageLimit() int {
	html := c.GetHtml()
	return utils.GetPageLimit(html)
}

// ReapSubjects Reap the subject content
func (c *Category) ReapSubjects(db *gorm.DB) []Subject {
	html := c.GetHtml()

	reg, _ := regexp.Compile(`<a href="(.*)" target="_blank" title="(.*)"`)
	dst := []byte("")
	template := "$1:$2"
	regColon, _ := regexp.Compile(`:`)
	for _, subj := range reg.FindAllString(html, -1) {
		var obj Subject

		if match, _ := regexp.MatchString(".xml", subj); !match {
			match := reg.FindStringSubmatchIndex(subj)
			tmp := regColon.Split(string(reg.ExpandString(dst, template, subj, match)), 2)
			obj.URL = tmp[0]
			obj.Title = tmp[1]
			obj.Images = obj.ReapImages(db)
			obj.CategoryID = c.ID
			//fmt.Printf("Image: %v\n", obj)
			c.Subjects = append(c.Subjects, obj)
			obj.Create(db)
		}
	}
	return c.Subjects
}
