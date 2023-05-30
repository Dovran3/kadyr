package src

import (
	"project/src/api"
	mw "project/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Init_app() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(gin.Logger())
	router.Use(mw.Cors) // it needs for fronted, to make request

	api.Controller(router)

	return router
}
