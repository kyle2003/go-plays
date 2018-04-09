package modules

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Theme struct {
	Name string
	Page int
}

//const base = "http://xxgege.net/"
const base = "https://studygolang.com/topics"

// ParseTheme to retrieve data
func ParseTheme(cate string) (topics []Topic, err error) {
	log.Println(cate)

	var t Topic
	url := base
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Failed to get content from %v", url)
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Failed to get content from %v", url)
		return
	}

	doc.Find(".topic").Each(func(i int, s *goquery.Selection) {
		t.Topic = s.Find(".title a").Text()
		s.Find(".meta a").Each(func(i int, ss *goquery.Selection) {
			if ss.HasClass("node") {
				t.Title = ss.Text()
				t.Thumb, _ = ss.Attr("href")
			}
		})

		topics = append(topics, t)
	})

	return
}

func BuildThemeResponse() {

}
