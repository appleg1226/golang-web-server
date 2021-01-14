package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUsers(c * gin.Context){
	var user []models.User
	err := models.GetAllUsers(&user)

	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c * gin.Context){
	var user models.User

	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateUser(&user); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)
}

func GetUserById(c * gin.Context){
	var user models.User
	id := c.Params.ByName("id")

	if err := models.GetUserById(&user, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c * gin.Context){
	var user models.User
	id := c.Params.ByName("id")

	if err := models.GetUserById(&user, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.UpdateUser(&user); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c * gin.Context){
	var user models.User
	id := c.Params.ByName("id")

	if err := models.DeleteUser(&user, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"id"+id: "is deleted"})
}