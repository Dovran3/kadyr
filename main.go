package main

import (
	
	db "admin/src/database"
	"admin/src"
   "github.com/gin-gonic/gin"
)

func main(){
	db.Init_db()
	router := gin.Default()
    
	router.GET("/users",src.GetUsers)
	//router.POST("/post",src.CeateUsers)
    router.POST("/add-data",src.PostData)
	
    router.Run(":8091")

}
