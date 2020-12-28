package controllers

import (
	"github.com/gin-gonic/gin"
)

//User type
type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

//UserGet return create
func UserGet(c *gin.Context) {
	//Receive token.id and return user

	user := User{
		Name: "Jorge",
		ID:   1,
	}

	tokenID := c.GetUint64("id")

	if tokenID != user.ID {
		c.JSON(200, gin.H{
			"Data": "Token Incorreto",
		})
	}

	// json.Marshal(&user)
	c.JSON(200, gin.H{
		"Data": &user,
	})
}
