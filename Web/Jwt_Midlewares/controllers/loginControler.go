package controllers

import (
	"strconv"

	"../middlewares"
	"github.com/gin-gonic/gin"
)

//Login return create
func Login(c *gin.Context) {

	user := User{
		Name: "Jorge",
		ID:   1,
	}

	idSend, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
	if user.ID != idSend {
		c.JSON(400, gin.H{
			"Data": "Login invalido",
		})
	} else {
		token := middlewares.GenerateJwt(strconv.FormatUint(user.ID, 10))

		message := ("tokem:" + token)
		c.JSON(200, gin.H{
			"Data": message,
		})
	}

}
