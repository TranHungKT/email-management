package listControllers

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

func CreateNewListController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var list models.List

		utils.BindJSON(ctx, &list)

		validationErr := validator.New().Struct(list)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := database.ListCollection().CountDocuments(context.TODO(), bson.D{primitive.E{Key: "name", Value: list.Name}})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "This list name is not available"})
			return
		}

		if list.Type == "" {
			list.Type = models.ListTypePrivate
		}

		if list.Optin == "" {
			list.Optin = models.ListOptinSingle
		}

		result, err := database.ListCollection().InsertOne(context.TODO(), &list)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusAccepted, result)
		ctx.Done()

	}

}
