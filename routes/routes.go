package routes

import "github.com/gin-gonic/gin"

func Useraroutes(incomingRoutes *gin.Engine) {
	
	incomingRoutes.POST("/users/signup")
	incomingRoutes.POST("/users/login")
	incomingRoutes.POST("/admin/addProduct")
	incomingRoutes.GET("/users/productView")
	incomingRoutes.GET("/users/search")
	
}