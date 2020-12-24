package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index return text hello
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": "Index Hello Controller"})
}
