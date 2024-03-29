package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"html/template"
	//"strings"

	globals "gin_session_auth/globals"
	middleware "gin_session_auth/middleware"
	routes "gin_session_auth/routes"
)

func main() {
	router := gin.Default()

	globals.ConnectDB()

	router.Static("/assets", "./assets")
	router.Static("/static/css", "./static/css")
	router.Static("/static/img", "./static/img")
	router.Static("/static/scss", "./static/scss")
	router.Static("/vendor", "./static/vendor")
	router.Static("/static/js", "./static/js")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))
	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("0.0.0.0:8080")
}
