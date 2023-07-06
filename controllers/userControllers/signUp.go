package userControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUpController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		utils.BindJSON(ctx, &user)

		validationErr := validator.New().Struct(&user)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := database.UserCollection().CountDocuments(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "This user already exist"})
			return
		}

		if user.Type == "" {
			user.Type = models.UserTypeUser
		}

		if user.Status == "" {
			user.Status = models.UserStatusEnabled
		}

		result, err := database.UserCollection().InsertOne(context.TODO(), &user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusAccepted, result)
		ctx.Done()
	}
}
