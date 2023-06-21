package cookie

import (
	"net/http"
	"time"
)

func NewCookieForJwtToken(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = false
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.HttpOnly = true
	return cookie
}
