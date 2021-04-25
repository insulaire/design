package handles

import (
	"design/internal/entities"
	"design/internal/result"
	"design/pkg/md5"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := entities.UrlPath{}
		c.BindJSON(&urlPath)
		if token, err := GetToken(c); err != nil {
			c.JSON(http.StatusOK, result.FailureResult(result.WithMsg(err.Error())))
			return
		} else {
			urlPath.WithToken(token)
		}
		url := md5.New(urlPath.Path)
		c.JSON(http.StatusOK, result.SueccResult(result.WithData(url)))
	}
}
