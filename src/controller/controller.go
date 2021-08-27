package controller

import (
	"mongo_go/src/database"
	"mongo_go/src/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("D:\\GO_Workspace\\src\\mongo_go\\templates\\*.html")

	//index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//GET ADD
	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add-user.html", nil)
	})
	// POST ADD
	r.POST("/add", routes.AddUser)

	// GET FIND
	r.GET("/find", func(c *gin.Context) {
		c.HTML(http.StatusOK, "find-user.html", nil)
	})
	// POST FIND
	r.POST("/find", routes.FindUser)

	// GET DELETE
	r.GET("/delete", func(c *gin.Context) {
		c.HTML(http.StatusOK, "delete-user.html", nil)
	})
	// POST DELETE
	r.POST("/delete", routes.DeleteUser)
	// TODO

	update := r.Group("/update")
	// GET UPDATE root to search for username
	update.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "update-find.html", nil)
	})
	// POST UPDATE root to search for username
	update.POST("/", func(c *gin.Context) {
		username := strings.TrimSpace(c.PostForm("username"))
		// c.String(http.StatusOK, "/update/edit/"+username)
		c.Redirect(http.StatusMovedPermanently, "/update/edit/"+username)
	})
	// GET UPDATE TO EDIT
	update.GET("/edit/:username", func(c *gin.Context) {
		username := c.Param("username")
		// c.String(http.StatusOK, username)
		result, err := database.FindUser(username)
		if err != nil {
			c.JSON(http.StatusNotFound, "No User Found")
			return
		}
		c.HTML(http.StatusOK, "update-user.html", result)
	})
	// POST UPDATE TO EDIT
	update.POST("/edit/", routes.UpdateUser)

	r.Static("/public", "D:\\GO_Workspace\\src\\mongo_go\\public")
	return r
}
