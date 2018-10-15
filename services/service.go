package services

import (
	"net/http"
	"pandora/conf"
	"pandora/constants"
	"pandora/models"
	"pandora/operations"
	"pandora/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// Start the craw and http service
func Start() {
	// Init category
	initCategory()
	if len(operations.FetchCategoryList()) > 0 {
		initSubject()
	}
	initDownload()
}

func init() {
	// Init glob db
	db := conf.GlobalDb.Get()
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Image{})
}

// initCategory
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
		if conf.Setup.Section("download").Haskey("default_limit") {
			c.Limit, _ = strconv.Atoi(conf.Setup.Section("download").Key("default_limit").String())
		}
		c.Create(db)
	}
}

// initSubject
func initSubject() {
	db := conf.GlobalDb.Get()
	cList := operations.FetchUnReapedCategoryList()
	for _, c := range cList {
		err := c.Reap(db)
		if err != nil {
			logrus.Warnln("Failed to reap subjects for category: [ " + c.Title + " ]")
		}
	}
}

// initDownload
func initDownload() {
	sList := operations.FetchReapedSubjectList()
	imgPath := conf.Setup.Section("download").Key("image_path").String()

	for _, s := range sList {
		cTitle := operations.GetCategoryTitleByID(s.CategoryID)
		err := utils.ProcessDir(imgPath + cTitle + "/" + s.Title)
		if err != nil {
			logrus.Warnf("%v", err)
			continue
		}
		operations.DownloadSubject(s.ID)
	}

	db := conf.GlobalDb.Get()
	for _, s := range sList {
		images := operations.GetNotDownloadedImagesBySubjectID(s.ID)
		if len(images) != 0 {
			s.DownloadStatus = constants.DOWNLOAD_STATUS__DONE
			db.Save(s)
		}
	}
}

// Run run web service
func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("index/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.Run()
}
