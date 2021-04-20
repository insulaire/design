package handles

import (
	"crypto/md5"
	"design/internal/entities"
	"design/internal/result"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := entities.UrlPath{}
		c.BindJSON(&urlPath)
		if token, err := GetToken(c); err != nil {
			c.JSON(http.StatusOK, result.FailureResult(result.ResultWithMsg(err.Error())))
			return
		} else {
			urlPath.WithToken(token)
		}

		hash := md5.New()
		io.WriteString(hash, urlPath.String())
		url := fmt.Sprintf("%x", hash.Sum([]byte{}))
		c.JSON(http.StatusOK, result.SueccResult(result.ResultWithData(url)))
	}
}
