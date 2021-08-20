package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("D:\\GO_Workspace\\src\\mongo_go\\templates*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Static("/public", "D:\\GO_Workspace\\src\\mongo_go\\public")
	return r
}
