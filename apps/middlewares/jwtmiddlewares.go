package middlewares

// import (
// 	_userDomain "final-project/business/users"
// 	"time"

// 	_jwt "github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// )

// func CreateToken(data _userDomain.Domain) (string, error) {
// 	claims := _jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userId"] = data.Id
// 	claims["exp"] = time.Now().Add(time.Hour * 6)
// 	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte("ThisisSecretGais"))
// }

// func extractToken(e echo.Context) uint {
// 	user := e.Get("user").(*_jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(_jwt.MapClaims)
// 		userId := claims["userId"].(uint)
// 		return userId
// 	}
// 	return 0;


// func ExtractTokenMetadata(c echo.Context) {
	// fmt.Println(c.Get("Authorization"))
	// bearer := c.Request().Header.Get("Authorization")
	// jwttoken := strings.Replace(bearer, "bearer ", "", 1)
	// claims := _jwt.MapClaims{}
	// _, err := jwt.ParseWithClaims(jwttoken, claims, func(jwttoken *_jwt.Token) (interface{}, error) {
	// 	return jwttoken.SignedString([]byte("ThisisSecretGais")), nil
	// })
	// if err != nil {
	// 	fmt.Println(claims)
	// }

	// ... error handling

	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

// }
// func Auth(c echo.Context) (bool, error) {

// 	return true, nil
// }