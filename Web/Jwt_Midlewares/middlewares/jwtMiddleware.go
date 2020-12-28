package middlewares

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	jwt "github.com/gbrlsnchs/jwt"
	"github.com/gin-gonic/gin"
)

//VerifyJwt verify jwt is valid
func VerifyJwt() gin.HandlerFunc {

	return func(c *gin.Context) {
		hs := jwt.NewHS256([]byte("secret"))
		pl := jwt.Payload{}

		HeaderToken := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		token := []byte(HeaderToken)

		hd, err := jwt.Verify(token, hs, &pl)
		if err != nil {
			log.Fatal(err)
		}
		tokenInt, _ := strconv.ParseUint(hd.KeyID, 10, 64)
		fmt.Println(tokenInt)
		c.Set("id", tokenInt)
		c.Next()
	}
}

//GenerateJwt return JWT
func GenerateJwt(ID string) string {

	hs := jwt.NewHS256([]byte("secret"))
	now := time.Now()
	pl := jwt.Payload{
		Issuer:   "gbrlsnchs",
		Subject:  "login",
		IssuedAt: jwt.NumericDate(now),
	}
	token, err := jwt.Sign(pl, hs, jwt.ContentType("JWT"), jwt.KeyID(ID))
	if err != nil {
		log.Fatal(err)
	}
	return string(token)

	//uint8
}
