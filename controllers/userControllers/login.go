package userControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/middleware"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.UserBase

		utils.BindJSON(ctx, &user)
		validationErr := validator.New().Struct(&user)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var foundUser models.User
		err := database.UserCollection().FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&foundUser)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if foundUser.Password != user.Password {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Email or Password"})
			return
		}

		middleware.HandleToken(ctx, foundUser)
		ctx.JSON(http.StatusNoContent, "")
		ctx.Done()
	}

}
