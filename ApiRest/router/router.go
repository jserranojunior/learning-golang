package router

import (
	books "../books/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Routes return all Routes
func Routes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.Run()
	//Books
	RouterBooks := router.Group("/books")
	{
		RouterBooks.GET("/", books.Index)
		RouterBooks.GET("/:id", books.Edit)
		RouterBooks.POST("/", books.Store)
		RouterBooks.PUT("/", books.Update)
		RouterBooks.DELETE("/", books.Delete)
	}

	//Return
	return router
}
