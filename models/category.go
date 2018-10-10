package models

import (
	"pandora/constants"
	"pandora/modules/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	// object的属性
	PandoraObj
	// 涉及范围
	Limit int
	// 主题数
	SubjectsNum int
}

func (c *Category) Create(db *gorm.DB) error {
	c.Name = "test"
	c.Title = "teset"
	c.URL = "http://test"
	c.ReapStatus = constants.REAP_STATUS__NOTDONE
	c.DownloadStatus = constants.DOWNLOAD_STATUS__NOTDONE
	c.Created = time.Now().Unix()
	c.Updated = time.Now().Unix()
	c.Enabled = uint8(1)
	c.SubjectsNum = 1
	c.Limit = 2
	err := db.Create(c).Error

	return err

}

func (th Category) GetHtml() string {
	url := constants.BASE + th.Title
	return string(utils.GetHtml(url))
}

func (th Category) GetPageLimit() int {
	html := th.GetHtml()
	return utils.GetPageLimit(html)
}

/*
func (th Category) GetSubjects() []subject.Subject {
	html := th.GetHtml()

	reg, _ := regexp.Compile(`<a href="(.*)" target="_blank" title="(.*)"`)
	dst := []byte("")
	template := "$1:$2"
	regColon, _ := regexp.Compile(`:`)
	var subjects []subject.Subject
	for _, subj := range reg.FindAllString(html, -1) {
		var obj subject.Subject

		if match, _ := regexp.MatchString(".xml", subj); !match {
			match := reg.FindStringSubmatchIndex(subj)
			tmp := regColon.Split(string(reg.ExpandString(dst, template, subj, match)), 2)
			obj.SubHref = tmp[0]
			obj.SubTitle = tmp[1]
			obj.Images = obj.GetImages()
			//fmt.Printf("Image: %v\n", obj)
			subjects = append(subjects, obj)
		}
	}

	return subjects
}
*/
