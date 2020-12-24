package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Store return create
func Store(c *gin.Context) {
	c.JSON(407, gin.H{
		"Data": "Create Books Controller",
	})
}

//Index return read
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Read Books Controller"})
}

//Update return Update
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Update Books Controller"})
}

//Delete return text delete
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Delete Books Controller"})
}
