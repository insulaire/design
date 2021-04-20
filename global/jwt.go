package global

import (
	"time"
)

type JWT struct {
	Exp    time.Duration
	Secret string
}

var GlbJWT JWT

func init() {
	GlbJWT = JWT{
		Exp: time.Hour,
		//Secret: uuid.New().String(),
		Secret: "demo",
	}
}
