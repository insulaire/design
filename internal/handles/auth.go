package handles

import (
	"design/pkg/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		str, err := jwt_auth.GenerateToken(map[string]interface{}{"user": "abc"})
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": str})
	}
}
