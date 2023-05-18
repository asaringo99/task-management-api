package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	UserId int  `json:"userid"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func CreateJwtToken(userid int, isAdmin bool) (string, error) {
	claims := &JwtCustomClaims{
		userid,
		isAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func RetrieveUserIdFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	userid := claims.UserId
	return userid
}
