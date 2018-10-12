package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pandora/constants"
	"pandora/utils"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Subject struct {
	// generic attributes
	PandoraObj
	// categoryID where subject belongs to
	CategoryID uint64 `gorm:"column:f_category_id;index;" json:""`
	// images collection num
	ImagesNum int8 `gorm:"column:f_images_num;default:0;" json:""`
	// the thumb imageid
	ThumbImageID uint64 `gorm:"column:f_thumb_image_id;default:0;" json:""`
	// images object collection
	Images []*Image `gorm:"-" json:"-"`
}

// Create db
func (s *Subject) Create(db *gorm.DB) error {
	// default attributes
	s.Created = time.Now().Unix()
	s.Updated = time.Now().Unix()

	db.AutoMigrate(s)
	return db.Create(s).Error
}

func (s *Subject) ReapImages(db *gorm.DB) error {
	url := s.URL
	html := utils.GetHtml(url)
	reg, _ := regexp.Compile(`img src="//(.*.jpg)"`)
	urlsStr := reg.FindString(html)
	reg, _ = regexp.Compile(`img|alt`)
	regImg, _ := regexp.Compile(`.*//(.*.jpg).*`)
	repl := "${1}"

	// 截取到的字符串
	for _, str := range reg.Split(urlsStr, -1) {
		if match, _ := regexp.MatchString(".*jpg", str); match {
			var img *Image
			str = strings.Replace(str, " ", "", -1)
			url := "http://" + regImg.ReplaceAllString(str, repl)

			if m, _ := regexp.MatchString(".jpg$", url); m {
				img.Name = utils.Basename(str)
				img.URL = url
				img.Title = s.Title
				img.SubjectID = s.ID
				img.CategoryID = s.CategoryID
				img.ReapStatus = constants.REAP_STATUS__DONE
				s.Images = append(s.Images, img)
				img.Create(db)
				s.ImagesNum++
			}
		}
	}

	if s.ImagesNum == 0 {
		return errors.New("Reaped 0 image links for " + s.Title)
	}
	return nil
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
