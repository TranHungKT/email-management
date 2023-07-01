package router

import (
	"github.com/TranHungKT/email_management/controllers/userControllers"
	"github.com/TranHungKT/email_management/middleware"
	"github.com/adam-hanna/jwt-auth/jwt"
	"github.com/gin-gonic/gin"
)

var restrictedRoute jwt.Auth
var HMACKey []byte

func InitGin() {
	middleware.InitRestrictedRoute()

	var router = gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Restricted")
	})

	UserRoutes(router)
	router.Run()
}

func UserRoutes(router *gin.Engine) {
	router.POST("/sign-up", userControllers.SignUpController())
	router.POST("/login", userControllers.LoginController())
}
