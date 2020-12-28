package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Store return create
func Store(c *gin.Context) {

	message := ("Store Books Controller " + c.PostForm("name"))
	c.JSON(200, gin.H{
		"Data": message,
	})
}

//Index return read
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Read Books Controller"})
}

//Edit return read
func Edit(c *gin.Context) {
	id := c.Param("id")
	message := ("Read Books Controller " + id)
	c.JSON(http.StatusOK, gin.H{"data": message})
}

//Update return Update
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Update Books Controller"})
}

//Delete return text delete
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Delete Books Controller"})
}
