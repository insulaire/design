package middlewares

import (
	"design/internal/result"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Limiter struct {
	//key string
	limiters sync.Map //[string]*rate.Limiter
}

func NewLimiter() *Limiter {
	return &Limiter{
		limiters: sync.Map{},
	}
}
func (limit *Limiter) Add(key string) {
	limit.limiters.Store(key, rate.NewLimiter(rate.Every(time.Second), 1))
}

var GlbLimiter *Limiter

func Init() {
	GlbLimiter = NewLimiter()
}

// NewMiddleware return a new instance of a gin middleware.
func NewMiddleware(limiter *Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter != nil {
			if v, ok := limiter.limiters.Load(c.Request.URL.String()); ok {
				limit, _ := v.(*rate.Limiter)
				//ctx, _ := context.WithTimeout(context.Background(), time.Second)
				//if err := limit.Wait(ctx); err != nil {
				if ok := limit.Allow(); ok {
					//c.JSON(http.StatusOK, result.FailureResult(result.WithMsg(err.Error())))
					c.JSON(http.StatusOK, result.FailureResult(result.WithMsg("so fast")))
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
