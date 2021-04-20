package middlewares

import (
	"design/pkg/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GWT(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	_, err := jwt_auth.ValidToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	c.Keys = map[string]interface{}{}
	c.Keys["token"] = token
	c.Next()
}
