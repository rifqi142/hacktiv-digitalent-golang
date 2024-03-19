package controllers

import (
	"fmt"
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/database"
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/helpers"
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	fmt.Println(db)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := &models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"email": User.Email,
		"full_name": User.FullName,
	})
}