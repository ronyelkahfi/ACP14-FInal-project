package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const jwtSecret = "SECRETges%&*^&*^*&("

type JwtCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(id uint, name string) (string, error) {
	claims := &JwtCustomClaims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
func Auth(c *echo.Context) bool {

	tokenString := echo.HeaderAuthorization
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("thisissecret"), nil
	})

	if token != nil && err == nil {
		return true
	} else {
		return false
	}
}

func ExtractClaims(ctx echo.Context) (JwtCustomClaims, error) {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return JwtCustomClaims{
		Id:   claims.Id,
		Name: claims.Name,
	}, nil
}
