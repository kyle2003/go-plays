package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pandora/constants"
	"pandora/modules/utils"
	"regexp"
	"strings"
)

type Subject struct {
	PandoraObj
	CategoryID uint64
	ImagesNum  int8
}

func (sub Subject) GetImages() []Image {
	url := constants.BASE + sub.URL
	html := utils.GetHtml(url)
	reg, _ := regexp.Compile(`img src="//(.*.jpg)"`)
	urlsStr := reg.FindString(string(html))
	reg, _ = regexp.Compile(`img|alt`)
	regImg, _ := regexp.Compile(`.*//(.*.jpg).*`)
	repl := "${1}"
	var images []Image

	// 截取到的字符串
	for _, str := range reg.Split(urlsStr, -1) {
		if match, _ := regexp.MatchString(".*jpg", str); match {
			var img Image
			str = strings.Replace(str, " ", "", -1)
			url := "http://" + regImg.ReplaceAllString(str, repl)

			if m, _ := regexp.MatchString(".jpg$", url); m {
				img.Name = utils.Basename(str)
				img.URL = url
				img.Title = sub.Title
			}

			images = append(images, img)
		}
	}
	return images
}

func DownloadImg(img Image) {
	resp, err := http.Get(img.URL)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Buidl path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := "图片/" + img.Title + "/" + img.Name + ".jpg"
	fh, err = os.Create(file)
	defer fh.Close()
	if err != nil {
		log.Fatalf("Failed to create img file: %s", file)
	}

	fh.Write(imgByte)
}
