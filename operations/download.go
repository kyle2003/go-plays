package operations

import (
	"io/ioutil"
	"net/http"
	"os"
	"pandora/models"

	"github.com/sirupsen/logrus"
)

func Download(img *models.Image) {
	resp, err := http.Get(img.URL)
	if err != nil {
		logrus.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Build path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := "图片/" + img.Title + "/" + img.Name + ".jpg"
	fh, err = os.Create(file)
	if err != nil {
		logrus.Fatalf("Failed to create img file: %s", file)
	} else {
		logrus.Printf("Creating: %s", file)
	}

	defer fh.Close()
	fh.Write(imgByte)
}
