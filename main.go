package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pandora/constants"
	"pandora/models"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
	var th = theme.Theme{
		TheTitle: "artzp",
	}
	utils.ProcessDir("图片")

	for _, sub := range th.GetSubjects() {
		utils.ProcessDir("图片/" + sub.SubTitle)

		for _, img := range sub.Images {
			fmt.Printf("Downloading: %v\n", img.ImgHref)
			//Download(img)

			//time.Sleep(time.Duration(15) * time.Second)
		}
	}
*/
func main() {
	db, err := gorm.Open("sqlite3", "./db/test.db")
	if err != nil {
		return
	}
	defer db.Close()
	c := models.Category{}
	c.Name = "test"
	c.Title = "teset"
	c.URL = "http://test"
	c.ReapStatus = constants.REAP_STATUS__NOTDONE
	c.DownloadStatus = constants.DOWNLOAD_STATUS__NOTDONE
	c.Created = time.Now().Unix()
	c.Updated = time.Now().Unix()
	c.Enabled = uint8(1)
	c.SubjectsNum = 1
	c.Limit = 2

	db.AutoMigrate(&c)
	db.Create(&c)
}

func Download(img models.Image) {
	resp, err := http.Get(img.URL)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Build path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := "图片/" + img.Title + "/" + img.Name + ".jpg"
	fh, err = os.Create(file)
	if err != nil {
		log.Fatalf("Failed to create img file: %s", file)
	} else {
		log.Printf("Creating: %s", file)
	}

	defer fh.Close()

	fh.Write(imgByte)
}
