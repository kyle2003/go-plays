package modules

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Topic struct {
	Topic string `json:"topic"`
	Title string `json:"title"`
	Thumb string `json:"thumb"`
}

type TopicDetails struct {
	title string
	imgs  []string
}

// ParseTopic to return topic details
func ParseTopic(t string) (data interface{}, err error) {
	url := base + t
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

func BuildTopicResponse() {

}
