package router

import (
	. "data-collector/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", ListUsers)
	router.GET("/user/:id", ListUser)
	router.POST("/user", CreateUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)

	return router
}
