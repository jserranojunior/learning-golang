package router

import (
	"../controllers"
	"../middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Routes return all Routes
func Routes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.Run()
	//Books
	RouterLogin := router.Group("/login")
	{
		RouterLogin.POST("/", controllers.Login)
	}
	RouterUser := router.Group("/user")
	{
		RouterUser.GET("/", middlewares.VerifyJwt(), controllers.UserGet)
	}

	//Return
	return router
}
