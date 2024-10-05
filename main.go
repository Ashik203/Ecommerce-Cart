package main

import (
	"log"
	"os"

	"github.com/Ashik203/ecommerce-cart/controllers"
	"github.com/Ashik203/ecommerce-cart/database"
	"github.com/Ashik203/ecommerce-cart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port ==""{
		port="8000"
	}

	app:= controllers.NewApplication(database.ProductData(database.Client,"Products"),database.UserData(database.Client,"Users"))

router:=gin.New()
router.Use(gin.Logger())

routes.UseraRoutes(router)
routes.Use(middleware.Authentication())

router.GET("/addToCart",app.AddToCart())
router.GET("/removeItem",app.RemoveItem())
router.GET("/cartCheckOut",app.BuyFromCart())
router.GET("/instantBuy",app.InstantBuy())

log.Fatal(router.Run(":"+ port))


}