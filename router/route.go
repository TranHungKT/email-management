package router

import (
	"github.com/TranHungKT/email_management/controllers/listControllers"
	"github.com/TranHungKT/email_management/controllers/subscriberControllers"
	"github.com/TranHungKT/email_management/controllers/userControllers"
	"github.com/TranHungKT/email_management/middleware"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	middleware.InitRestrictedRoute()

	var router = gin.Default()
	HTMLRender(router)

	UserRoutes(router)
	ListRoutes(router)
	SubscriberRoutes(router)
	router.Run()
}

func HTMLRender(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./static/public/*.html")
	router.GET("/subscriber/confirm-optin/:nonceKey/:cipherEmailKey/:startedTime", func(ctx *gin.Context) {
		nonce := ctx.Param("nonceKey")
		cipherEmail := ctx.Param("cipherEmailKey")

		email := utils.DecryptCipher(nonce, cipherEmail)

		ctx.HTML(200, "confirmOptin.html", map[string]string{"email": email})
	})

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
	router.POST("/subscriber/confirm-optin/:nonceKey/:cipherEmailKey/:startedTime", subscriberControllers.ConfirmOptinController())
}
