package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pandora/constants"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"

	"pandora/database"
	"pandora/models"
	"pandora/utils"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	dbObj := &database.SqliteObj{}
	db := dbObj.Get()

	c := &models.Category{
		SubjectsNum: 1,
	}
	c.Name = "artzp"
	c.Title = "自拍"
	c.URL = constants.BASE + c.Name
	c.Create(db)
	utils.ProcessDir("图片")

	for _, sub := range c.ReapSubjects(db) {
		utils.ProcessDir("图片/" + sub.SubTitle)

		for _, img := range sub.Images {
			fmt.Printf("Downloading: %v\n", img.URL)
			//Download(img)

			//time.Sleep(time.Duration(15) * time.Second)
		}
	}
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
