package operations

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"pandora/conf"
	"pandora/constants"
	"pandora/models"

	"github.com/sirupsen/logrus"
)

func DownloadSubject(s *models.Subject) {
	for _, i := range s.Images {
		go Download(i)
	}
}

func Download(img *models.Image) {
	resp, err := http.Get(img.URL)
	if err != nil {
		logrus.Printf("%v", err)
	}
	defer resp.Body.Close()

	// Build path
	imgByte, err := ioutil.ReadAll(resp.Body)

	var fh *os.File
	file := conf.Setup.Section("download").Key("image_path").String() + img.Title + "/" + img.Name + ".jpg"
	fh, err = os.Create(file)
	if err != nil {
		logrus.Fatalf("Failed to create img file: %s", file)
	} else {
		logrus.Printf("Creating: %s", file)
	}

	defer fh.Close()
	fh.Write(imgByte)

	// Save base64 to db
	db := conf.GlobalDb.Get()

	img.Base64 = base64.StdEncoding.EncodeToString(imgByte)
	img.DownloadStatus = constants.DOWNLOAD_STATUS__DONE
	db.Save(img)
}
