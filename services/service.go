package services

import (
	"fmt"
	"pandora/conf"
	"pandora/constants"
	"pandora/models"
	"pandora/operations"
	"pandora/utils"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// Start the craw and http service
func Start() {
	// Init category
	initCategory()
	initSubject()
	//logrus.Debugf("%v", operations.FetchCategoryList())
}

func init() {
	// Init glob db
}

func initCategory() {
	// load category picked from setup.ini file
	cfg, err := ini.Load("conf/setup.ini")
	if err != nil {
		logrus.Fatalf("Failed to load conf file: %v", err)
	}

	categories := cfg.Section("category").KeysHash()
	if len(categories) == 0 {
		logrus.Warn("The category were not set on conf: conf/setup.ini")
		return
	}

	cList := operations.FetchCategoryList()
	for _, c := range cList {
		if _, ok := categories[c.Name]; ok {
			delete(categories, c.Name)
		}
	}

	db := conf.GlobalDb.Get()
	for name, title := range categories {
		c := &models.Category{}
		c.Name = name
		c.Title = title
		c.URL = constants.BASE + "/" + c.Name + "/"
		c.Create(db)
	}
}

// initSubject
func initSubject() {
	db := conf.GlobalDb.Get()
	cList := operations.FetchCategoryList()
	imgPath := conf.Setup.Section("download").Key("image_path").String()
	for _, c := range cList {
		err := c.ReapSubjects(db)
		if err == nil {
			for _, s := range c.Subjects {
				err := utils.ProcessDir(imgPath + c.Title + "/" + s.Title)
				if err != nil {
					logrus.Printf("Create dir failed: %v", err)
					break
				}
				for _, i := range s.Images {
					fmt.Printf("Downloading: %v\n", i.URL)
				}
				// update thumb image id
				s.ThumbImageID = operations.FetchThumbImageBySubjectID(s.ID)
				//logrus.Printf("thumbID: %v", s.ThumbImageID)
				db.Save(&s)
			}
		} else {
			logrus.Warnln("Failed to reap subjects for category: [ " + c.Title + " ]")
		}
	}
}

func initDownload() {

}
