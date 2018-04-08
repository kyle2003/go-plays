package modules

import (
	"fmt"
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

func (t *Theme) Handler(res http.ResponseWriter, req *http.Request) {
	data, err := GetTopicData(t.Name)

	fmt.Fprintf(res, "hello %q", t.Name)
	fmt.Fprintf(res, "hello %q", data)

	if err != nil {
		log.Fatal("Error geting topic")
	}

}

// GetTopicData Return the topic data
func GetTopicData(name string) (data []interface{}, err error) {
	data, err = ParseTheme(name)

	if err != nil {
		log.Fatal("Failed to retreive data from theme")
	}
	return
}

func ParseTopic(topic string) {

}

// ParseTheme to retrieve data
func ParseTheme(cate string) (topics []interface{}, err error) {
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
