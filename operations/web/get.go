package web

import (
	"net/http"
	"pandora/operations/inner"
	"strconv"

	"github.com/gin-gonic/gin"
)

// atoi
func atoi(c *gin.Context) uint64 {
	x, _ := strconv.Atoi(c.Param("cid"))
	return uint64(x)
}

// GetCategoryDetails Provide the detailed page who listing subjects
func GetCategoryDetails(c *gin.Context) {
	id := atoi(c)

	subjects := inner.FetchSubjectsByCategoryID(id)
	category := inner.GetCategoryByID(id)
	categories := inner.FetchCategoryList()

	c.HTML(http.StatusOK, "category.tmpl", gin.H{
		"Category":   category,
		"Categories": categories,
		"Subjects":   subjects,
	})
}

func GetSubjectDetails(c *gin.Context) {
	//cid, _ := strconv.Atoi(c.Param("cid"))
	x, _ := strconv.Atoi(c.Param("sid"))
	sid := uint64(x)

	categories := inner.FetchCategoryList()
	subjects := inner.GetSubjectByID(sid)
	images := inner.GetDownloadedImagesBySubjectID(sid)

	c.HTML(http.StatusOK, "subject.tmpl", gin.H{
		"Categories": categories,
		"Subjects":   subjects,
		"Images":     images,
	})
}
