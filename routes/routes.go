package routes

import (
	"/controllers/controllers"

	"github.com/gin-gonic/gin"
)

func Useraroutes(incomingRoutes *gin.Engine) {
	
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addProduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productView", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	
}