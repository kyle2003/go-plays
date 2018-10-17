package operations

import (
	"net/http"

	"pandora/operations/inner"
	"pandora/operations/web"

	"github.com/gin-gonic/gin"
)

// Start the web services
func Start() {
	categories := inner.FetchCategoryList()
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Static("/images", "./images")
	router.GET("index/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Categories": categories,
		})
	})
	router.GET("index/:cid", web.GetCategoryDetails)
	router.GET("index/:cid/:sid", web.GetSubjectDetails)
	router.Run()
}
