package models

import (
	"pandora/constants"
	"pandora/models/subject"
	"pandora/modules/utils"
	"regexp"
)

type Category struct {
	// 自增ID
	ID uint
	// 拼音名
	Name string
	// 中文名
	Title string
	// URL地址
	URL string
	// 涉及范围
	Range string
	// 是否采集完成
	Reaped constants.ReapStatus
	// 主题数
	Subjects int
	// 创建时间戳
	Created uint64
	// 更新时间戳
	Updated uint64
}

func (th Category) GetHtml() string {
	url := constants.BASE + th.TheTitle
	return string(utils.GetHtml(url))
}

func (th Category) GetPageLimit() int {
	html := th.GetHtml()
	return utils.GetPageLimit(html)
}

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
