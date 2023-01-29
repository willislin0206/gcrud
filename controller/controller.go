package controller

import (
	model "data-collector/model"
	service "data-collector/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// "strconv"
)

//list repository
func ListUsers(c *gin.Context) {

	result, err := service.List()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//list someuser
func ListUser(c *gin.Context) {

	id := c.Param("id")
	result, err := service.ListUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//add user
func CreateUser(c *gin.Context) {
	var user model.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Error(err)
		return
	}

	err = service.Insert(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "added failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "added sucessfully",
		"data":    user.ID,
	})
}

//update user record
func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Error(err)
		return
	}

	err = service.UpdateUser(id, &user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "updated failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "updated sucessfully",
		"data":    user,
	})
}

//delete user
func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	log.Info(id)

	_, err := service.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "deleted failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "deleted sucessfully",
		"data":    id,
	})
}
