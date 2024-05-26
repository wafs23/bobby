package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"bobby/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.New()

	router.Use(gin.Recovery(), cors.Default())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		if param.ErrorMessage != "" {
			return fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | %s \n\n\t%s\n\n",
				param.TimeStamp.Format(time.RFC1123),
				param.StatusCodeColor(),
				param.StatusCode,
				param.ResetColor(),
				param.ClientIP,
				param.MethodColor(),
				param.Method,
				param.ResetColor(),
				param.Path,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		} else {
			return fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | %s\n",
				param.TimeStamp.Format(time.RFC1123),
				param.StatusCodeColor(),
				param.StatusCode,
				param.ResetColor(),
				param.ClientIP,
				param.MethodColor(),
				param.Method,
				param.ResetColor(),
				param.Path,
				param.Latency,
				param.Request.UserAgent(),
			)
		}
	}))

	store := cookie.NewStore([]byte(server.config.CookieKey))
	router.Use(sessions.Sessions("__sesbobby", store))

	/*
		CSRF
	*/
	// router.Use(middleware.SkipCheck())
	csrfSecure, _ := strconv.ParseBool(server.config.Secure)
	csrfMiddleware := csrf.Protect(
		[]byte(server.config.CsrfKey),
		csrf.Secure(csrfSecure),
		csrf.SameSite(csrf.SameSiteLaxMode),
		csrf.Path("/"),
	)
	router.Use(adapter.Wrap(csrfMiddleware))

	router.Static("/assets", "./assets")
	router.SetFuncMap(template.FuncMap{
		"add": util.Add,
	})
	router.LoadHTMLGlob("templates/*/*.html")

	router.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "🎈Page not found",
			},
		)
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(
			http.StatusMethodNotAllowed,
			gin.H{
				"message": "🎈Method not allowed",
			},
		)
	})

	/*
		PAGES
	*/
	router.GET("/", server.renderIndex)

	server.router = router
}

func (server *Server) Start() error {
	p := os.Getenv("PORT")
	if p == "" {
		p = "9033"
	}

	return server.router.Run(":" + p)
}
