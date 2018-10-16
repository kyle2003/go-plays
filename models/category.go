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

// Reap Reap the subject content
func (c *Category) Reap(db *gorm.DB) error {
	html := c.GetHTML(c.Limit)

	reg, _ := regexp.Compile(`<a href="(.*)" target="_blank" title="(.*)"`)
	dst := []byte("")
	template := "$1:$2"
	regColon, _ := regexp.Compile(`:`)
	for _, subjStr := range reg.FindAllString(html, -1) {
		var newSubj Subject

		if match, _ := regexp.MatchString(".xml", subjStr); !match {
			match := reg.FindStringSubmatchIndex(subjStr)
			tmp := regColon.Split(string(reg.ExpandString(dst, template, subjStr, match)), 2)
			newSubj.Name = tmp[0]
			newSubj.URL = constants.BASE + tmp[0]
			newSubj.Title = tmp[1]
			newSubj.CategoryID = c.ID

			// If subject existed
			db.Where(&newSubj).First(&newSubj)
			if newSubj.ID == uint64(0) {
				newSubj.Create(db)
				c.SubjectsNum++
			}

			// Check if images reaped
			if newSubj.ReapStatus == constants.REAP_STATUS__NOTDONE {
				err := newSubj.Reap(db)
				if newSubj.ImagesNum == 0 {
					newSubj.ReapStatus = constants.REAP_STATUS__NOTDONE
					newSubj.Enabled = constants.BOOL__FALSE
				} else {
					newSubj.ReapStatus = constants.REAP_STATUS__DONE
				}
				db.Save(&newSubj)

				if err != nil {
					logrus.Warningf("%v", err)
					continue
				}
			}
			//c.Subjects = append(c.Subjects, newSubj)
		}
	}

	if c.SubjectsNum == 0 {
		return errors.New("Reap 0 subjects for category " + c.Title)
	}
	c.ReapStatus = constants.REAP_STATUS__DONE
	db.Save(c)

	return nil
}
