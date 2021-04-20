package handles

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) (string, error) {
	if token, ok := c.Keys["token"]; !ok {
		return "", errors.New("token not found")
	} else {
		return token.(string), nil
	}
}
