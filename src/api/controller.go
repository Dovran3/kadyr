package api

import "github.com/gin-gonic/gin"

func Controller(router *gin.Engine) {
	router.GET("/users", get_users)
	router.POST("/new-user", create_user)
}
