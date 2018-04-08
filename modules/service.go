package modules

import (
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type topic struct {
	Topic string `json:"topic"`
	Title string `json:"title"`
	Thumb string `json:"thumb`
}

type Theme struct {
	Name string
}

//const base = "http://xxgege.net/"
const Base = "http://baidu.com"

func (t *Theme) Handler(res http.ResponseWriter, req *http.Request) {
	data, err := GetTopicData(t.Name)

	if err != nil {
		log.Fatal("Error geting topic")
	}
	//io.Copy(res, data)
	log.Println(data)
}

func GetTopicData(name string) (data []topic, err error) {
	data, err = ParseTheme(name)

	if err != nil {
		log.Fatal("Failed to retreive data from theme")
	}
	return
}

func ParseTheme(cate string) (topics []topic, err error) {
	url := base + cate
	resp, err := http.Get(url)
	log.Println(url)

	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Failed to get content from %v", url)
	}

	_, err = io.Copy(os.Stdout, resp.Body)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("Failed to get content from %v", url)
	}

	log.Println(doc.Attr, doc.Data)

	log.Println(doc)

	return
}
