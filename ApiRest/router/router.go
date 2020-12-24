package router

import (
	books "../books/controllers"
	hello "../hello/controllers"
	"github.com/gin-gonic/gin"
)

//Routes return all Routes
func Routes() *gin.Engine {
	r := gin.Default()

	//Books
	r.GET("/books", books.Index)
	r.GET("/books/store", books.Store)

	//Hello
	r.GET("/", hello.Index)

	//Return
	return r
}
