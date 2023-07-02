package userControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
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

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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

		result, err := database.UserCollection().InsertOne(context.TODO(), bson.D{
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
			{Key: "first_name", Value: user.FirstName},
			{Key: "last_name", Value: user.LastName},
			{Key: "status", Value: models.UserStatusEnabled},
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusAccepted, result)
		ctx.Done()
	}
}
