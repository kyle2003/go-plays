package web

import (
	"net/http"
	"pandora/operations/inner"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategoryDetails Provide the detailed page who listing subjects
func GetCategoryDetails(c *gin.Context) {
	x, _ := strconv.Atoi(c.Param("id"))
	id := uint64(x)

	subjects := inner.FetchSubjectsByCategoryID(id)
	category := inner.GetCategoryByID(id)

	c.HTML(http.StatusOK, "category.tmpl", gin.H{
		"Category": category,
		"Subjects": subjects,
	})
}
