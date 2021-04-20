package jwt_auth

import (
	"design/global"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var method *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

func GenerateToken(claims map[string]interface{}) (string, error) {
	if _, ok := claims["Exp"]; !ok {
		claims["Exp"] = time.Now().Add(global.GlbJWT.Exp).Unix()
	}

	c := jwt.NewWithClaims(method, jwt.MapClaims(claims))
	return c.SignedString([]byte(global.GlbJWT.Secret))
}

func ValidToken(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.GlbJWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	val := map[string]interface{}(t.Claims.(jwt.MapClaims))
	if _, ok := val["Exp"]; !ok {
		return nil, errors.New("Exp")
	}
	if exp, ok := val["Exp"].(float64); ok && exp < float64(time.Now().Unix()) {
		return nil, errors.New("Exp")
	}
	delete(val, "Exp")
	return val, err
}
