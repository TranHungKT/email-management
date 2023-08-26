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
	router.LoadHTMLGlob("./static/public/*")

	router.GET("/upload", func(ctx *gin.Context) {
		ctx.HTML(200, "upload.html", map[string]string{"title": "Ok"})
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

}
