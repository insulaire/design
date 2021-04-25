package Context

import (
	"design/internal/entities"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	user entities.User
}

// func (ctx *Context) GetUser() entities.User {

// }
