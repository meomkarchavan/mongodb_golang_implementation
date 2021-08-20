package routes

import (
	"mongo_go/src/database"
	"mongo_go/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.CreateUser(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}

func FindUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.FindUser(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeleteUser(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.UpdateUser(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}
