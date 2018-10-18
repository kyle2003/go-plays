package services

import (
	"pandora/conf"
	"pandora/constants"
	"pandora/models"
	"pandora/operations"
	"pandora/operations/inner"
	"pandora/utils"
	"strconv"
	"time"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// Start the craw and http service
func Start() {
	// Init category and harvest subjects
	go func() {
		initCategory()
		initSubject()
	}()

	go func() {
		// Need time to harvest the subjects
		for {
			time.Sleep(time.Duration(60) * time.Second)
			initDownload()
		}
	}()

	// Provide the web services`
	operations.Start()
}

// populate tables
func init() {
	// Init glob db
	db := conf.GlobalDb.Get()
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Image{})
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

	db := conf.GlobalDb.Get()
	for name, title := range categories {
		c := models.Category{}
		c.Name = name
		c.Title = title
		c.URL = constants.BASE + "/" + c.Name + "/"
		db.Where(&c).First(&c)
		if c.ID == uint64(0) {
			c.Create(db)
			db.Where(&c).First(&c)
		}
		if conf.Setup.Section("download").Haskey("default_limit") {
			c.Limit, _ = strconv.Atoi(conf.Setup.Section("download").Key("default_limit").String())
		}
		db.Save(&c)
	}
}

// initSubject
func initSubject() {
	db := conf.GlobalDb.Get()
	cList := inner.FetchUnReapedCategoryList()
	for _, c := range cList {
		if c.ReapStatus != constants.REAP_STATUS__DONE {
			err := c.Reap(db)
			if err != nil {
				logrus.Warnln("Failed to reap subjects for category: [ " + c.Title + " ]")
			}
		}
	}
}

// init Download
func initDownload() {
	sList := inner.FetchReapedSubjectList()
	imgPath := conf.Setup.Section("download").Key("image_path").String()

	for _, s := range sList {
		if s.DownloadStatus != constants.DOWNLOAD_STATUS__DONE {
			cTitle := inner.GetCategoryTitleByID(s.CategoryID)
			err := utils.ProcessDir(imgPath + cTitle + "/" + s.Title)
			if err != nil {
				logrus.Warnf("%v", err)
				continue
			}
			inner.DownloadSubject(&s)
		}
	}

	db := conf.GlobalDb.Get()
	for _, s := range sList {
		images := inner.GetNotDownloadedImagesBySubjectID(s.ID)
		if len(images) == 0 {
			s.DownloadStatus = constants.DOWNLOAD_STATUS__DONE
			db.Save(s)
		}
	}
}
