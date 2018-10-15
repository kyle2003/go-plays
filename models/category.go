package models

import (
	"errors"
	"pandora/constants"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"

	"regexp"
)

// Category struct
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

// Reap Reap the subject content
func (c *Category) Reap(db *gorm.DB) error {
	html := c.GetHTML(c.Limit)

	reg, _ := regexp.Compile(`<a href="(.*)" target="_blank" title="(.*)"`)
	dst := []byte("")
	template := "$1:$2"
	regColon, _ := regexp.Compile(`:`)
	for _, str := range reg.FindAllString(html, -1) {
		var newSubj Subject
		var existedSubj *Subject

		if match, _ := regexp.MatchString(".xml", str); !match {
			match := reg.FindStringSubmatchIndex(str)
			tmp := regColon.Split(string(reg.ExpandString(dst, template, str, match)), 2)
			newSubj.Name = tmp[0]
			newSubj.Title = tmp[1]
			newSubj.URL = constants.BASE + tmp[0]
			newSubj.CategoryID = c.ID

			// Check the subj existed already
			// if not create new one
			db.Where(&newSubj).First(existedSubj)
			if existedSubj == nil {
				newSubj.Create(db)
				c.SubjectsNum++
			}

			// Check the subj reaped
			if newSubj.ReapStatus != constants.REAP_STATUS__DONE {
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
	c.ReapStatus = constants.REAP_STATUS__DONE
	db.Save(c)

	if len(c.Subjects) == 0 {
		return errors.New("Reap 0 subjects for category " + c.Title)
	}
	return nil
}
