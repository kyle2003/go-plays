package modules

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type topic struct {
	Topic string `json:"topic"`
	Title string `json:"title"`
	Thumb string `json:"thumb"`
}

type Theme struct {
	Name string
}

//const base = "http://xxgege.net/"
const base = "https://studygolang.com/topics"

type TopicDetails struct {
	title string
	imgs  []string
}

// ParseTopic to return topic details
func ParseTopic(t topic) (data interface{}, err error) {
	url := base + t.Title
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
	doc.Find("test").Each(func(i int, s *goquery.Selection) {

	})
	return
}

// ParseTheme to retrieve data
func ParseTheme(cate string) (topics []topic, err error) {
	log.Println(cate)

	var t topic
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
