/*
=====================Golang======================
	Author 	: Audy Vaksi Pranata
	Link	: https://github.com/vaksi
	Email	: vaksipranata@gmail.com
==================================================
*/
package controllers

import(
	"go_auth_gin/models"

	"github.com/gin-gonic/gin"
)

var users models.Users

func UserAuth(username string, password string) (bool) {
	status := false
	for _, user := range users {
		if user.Username == username && user.Password == password {
			status = true
		}
	}
	return status
}

func GetUsers(c *gin.Context){
	c.JSON(200, gin.H{
		"status":  "OK",
		"users": users,
	})
}

func CreateUser(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	user.Username = username
	user.Password = password
	users = append(users, user)

	c.JSON(200, gin.H{
		"status":  "posted",
		"users": users,
	})
}