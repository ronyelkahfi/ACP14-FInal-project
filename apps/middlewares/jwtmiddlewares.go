package middlewares

import (
	_userDomain "final-project/business/users"
	"time"

	_jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(data _userDomain.Domain) (string, error) {
	claims := _jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = data.Id
	claims["exp"] = time.Now().Add(time.Hour * 6)
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("ThisisSecretGais"))
}

func ExtractToken(e echo.Context) uint {
	user := e.Get("user").(*_jwt.Token)
	if user.Valid {
		claims := user.Claims.(_jwt.MapClaims)
		userId := claims["userId"].(uint)
		return userId
	}
	return 0

}
