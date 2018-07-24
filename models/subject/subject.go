package subject

import (
	"fmt"
	"goplays/constants"
	"goplays/modules/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Image struct {
	ImgName     string
	ImgHref     string
	ImgSubTitle string
}

type Subject struct {
	SubTitle string
	SubHref  string
	Images   []Image
}

func (sub Subject) GetImages() []Image {
	url := constants.BASE + sub.SubHref
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
				img.ImgName = utils.Basename(str)
				img.ImgHref = url
				img.ImgSubTitle = sub.SubTitle
			}

			images = append(images, img)
		}
	}
	return images
}

func DownloadImg(img Image) {
	resp, err := http.Get(img.ImgHref)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Buidl path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := "图片/" + img.ImgSubTitle + "/" + img.ImgName + ".jpg"
	fh, err = os.Create(file)
	if err != nil {
		log.Fatalf("Failed to create img file: %s", file)
	}

	defer fh.Close()

	fh.Write(imgByte)
}
