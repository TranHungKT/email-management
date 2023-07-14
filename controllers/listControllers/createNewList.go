package listControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/controllers/handlers/listHandlers"
	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewListController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var list models.List

		err := utils.BindJSONAndValidateByStruct(ctx, &list)
		if err != nil {
			return
		}

		count, err := database.ListCollection().CountDocuments(context.TODO(), bson.D{primitive.E{Key: "name", Value: list.Name}})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "This list name is not available"})
			return
		}
		result := listHandlers.CreateNewListHandler(list)
		ctx.JSON(http.StatusAccepted, gin.H{"InsertedID": result})
		ctx.Done()
	}

}
