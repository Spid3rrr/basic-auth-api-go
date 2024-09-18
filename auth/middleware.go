package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatus(401)
		return
	}

	tokenString = strings.Split(tokenString, " ")[1]

	claims, err := VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	c.Set("username", claims["username"])
	c.Next()
}
