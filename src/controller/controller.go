package controller

import (
	"mongo_go/src/routes"
	"net/http"

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
	// GET UPDATE
	update.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "update-find.html", nil)
	})
	update.POST("/", func(c *gin.Context) {
		// username := strings.TrimSpace(c.PostForm("username"))
		c.Redirect(http.StatusOK, "/update/edit")
		// c.HTML(http.StatusOK, "update-user.html", username)
	})
	// GET UPDATE
	update.GET("/edit", func(c *gin.Context) {
		// c.String(http.StatusOK, c.Request.Body)

		// user, _ := database.FindUser(username)
		// c.HTML(http.StatusOK, "update-user.html", user)
	})
	// POST UPDATE
	update.POST("/update", routes.UpdateUser)

	r.Static("/public", "D:\\GO_Workspace\\src\\mongo_go\\public")
	return r
}
