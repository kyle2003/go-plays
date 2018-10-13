package models

import (
	"errors"
	"pandora/constants"

	"github.com/sirupsen/logrus"

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
	SubjectsNum int `gorm:"column:f_subjects_num;default:0;" json:"subjects_num"`
	// subjects
	Subjects []Subject `gorm:"-"`
}

// Create db
func (c *Category) Create(db *gorm.DB) error {
	// default attributes
	c.Created = time.Now().Unix()
	c.Updated = time.Now().Unix()

	db.AutoMigrate(c)
	return db.Create(c).Error
}

// ReapSubjects Reap the subject content
func (c *Category) ReapSubjects(db *gorm.DB) error {
	html := c.GetHtml(c.Limit)

	reg, _ := regexp.Compile(`<a href="(.*)" target="_blank" title="(.*)"`)
	dst := []byte("")
	template := "$1:$2"
	regColon, _ := regexp.Compile(`:`)
	for _, subj := range reg.FindAllString(html, -1) {
		var obj Subject

		if match, _ := regexp.MatchString(".xml", subj); !match {
			match := reg.FindStringSubmatchIndex(subj)
			tmp := regColon.Split(string(reg.ExpandString(dst, template, subj, match)), 2)
			obj.Name = tmp[0]
			obj.URL = constants.BASE + tmp[0]
			obj.Title = tmp[1]
			obj.CategoryID = c.ID
			err := obj.ReapImages(db)

			if err != nil {
				logrus.Warningf("%v", err)
				continue
			}
			logrus.Printf("thumb imgId: %v", obj.Images[0].ID)
			obj.ThumbImageID = obj.Images[0].ID
			obj.ReapStatus = constants.REAP_STATUS__DONE
			c.Subjects = append(c.Subjects, obj)
			obj.Create(db)
			c.SubjectsNum++
		}
		db.Save(c)
	}
	logrus.Println("test2")
	if len(c.Subjects) == 0 {
		return errors.New("Reap 0 subjects for category " + c.Title)
	}
	return nil
}
