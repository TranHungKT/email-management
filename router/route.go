package router

import (
	"github.com/TranHungKT/email_management/controllers/listControllers"
	"github.com/TranHungKT/email_management/controllers/subscriberControllers"
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
	SubscriberRoutes(router)
	router.Run()
}

func UserRoutes(router *gin.Engine) {
	router.POST("auth/sign-up", userControllers.SignUpController())
	router.POST("auth/login", userControllers.LoginController())
}

func ListRoutes(router *gin.Engine) {
	router.POST("list/create-new-list", middleware.RestrictedFunc(), listControllers.CreateNewListController())
}

func SubscriberRoutes(router *gin.Engine) {
	router.POST("/subscriber/create-new-subscriber", middleware.RestrictedFunc(), subscriberControllers.CreateNewSubscriberController())
	router.POST("/subscriber/confirm-optin", subscriberControllers.ConfirmOptinController())
	// router.GET("/subscriber/confirm-optin?email", html page here)

}
