package main

import (
	"fmt"
	"pandora/models/subject"
	"pandora/models/theme"
	"pandora/modules/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
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
}

func Download(img subject.Image) {
	resp, err := http.Get(img.ImgHref)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Build path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := "图片/" + img.ImgSubTitle + "/" + img.ImgName + ".jpg"
	fh, err = os.Create(file)
	if err != nil {
		log.Fatalf("Failed to create img file: %s", file)
	} else {
		log.Printf("Creating: %s", file)
	}

	defer fh.Close()

	fh.Write(imgByte)
}
