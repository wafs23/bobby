package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func (s *Server) customH(h gin.H, c *gin.Context) gin.H {

	h["csrf"] = csrf.TemplateField(c.Request)
	h["csrfv"] = csrf.Token(c.Request)
	h["thisUrl"] = fmt.Sprintf("%s://%s", s.config.URLScheme, s.config.Domain)

	return h
}
