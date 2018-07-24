package theme

import (
	"goplays/constants"
	"goplays/models/subject"
	"goplays/modules/utils"
	"regexp"
)

type Theme struct {
	TheTitle  string
	PageLimit int
	Subjects  []subject.Subject
}

func (th Theme) GetHtml() string {
	url := constants.BASE + th.TheTitle
	return string(utils.GetHtml(url))
}

func (th Theme) GetPageLimit() int {
	html := th.GetHtml()
	return utils.GetPageLimit(html)
}

func (th Theme) GetSubjects() []subject.Subject {
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
