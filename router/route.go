package router

import (
	"github.com/TranHungKT/email_management/controllers/listControllers"
	"github.com/TranHungKT/email_management/controllers/userControllers"
	"github.com/TranHungKT/email_management/middleware"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	middleware.InitRestrictedRoute()

	var router = gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Restricted")
	})

	UserRoutes(router)
	ListRoutes(router)
	router.Run()
}

func UserRoutes(router *gin.Engine) {
	router.POST("/sign-up", userControllers.SignUpController())
	router.POST("/login", userControllers.LoginController())
}

func ListRoutes(router *gin.Engine) {
	router.POST("/create-new-list", middleware.RestrictedFunc(), listControllers.CreateNewListController())
}
