package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) renderIndex(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"index.html",
		s.customH(
			gin.H{
				"title": "Hello M",
			},
			c,
		),
	)
}
