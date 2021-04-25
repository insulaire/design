package handles

import (
	"design/internal/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, result.SueccResult())
	}

}
