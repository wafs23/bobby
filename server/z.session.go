package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setSession(key string, data interface{}, secure bool, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set(key, data)
	session.Options(sessions.Options{
		MaxAge:   3600 * 24 * 7,
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
	err := session.Save()
	return err
}

func clearSession(key string, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set(key, "content")
	session.Options(sessions.Options{MaxAge: -1, Path: "/"})
	err := session.Save()
	return err
}
