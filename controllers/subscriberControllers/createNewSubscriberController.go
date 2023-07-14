package subscriberControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/controllers/handlers/listHandlers"
	"github.com/TranHungKT/email_management/controllers/handlers/subscriberHandlers"
	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewSubscriberController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newSubscriber models.NewSubscriberRequestPayload
		err := utils.BindJSONAndValidateByStruct(ctx, &newSubscriber)
		if err != nil {
			return
		}

		count, err := database.SubscriberCollection().CountDocuments(context.TODO(), bson.D{primitive.E{Key: "email", Value: newSubscriber.Email}})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "This list subscriber has already subscribed, please update list for it or use other email"})
			return
		}

		if newSubscriber.Status == "" {
			newSubscriber.Status = models.SubscriberStatusEnabled
		}

		lists, err := listHandlers.GetListByIdsHandler(newSubscriber.ListIds)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := subscriberHandlers.CreateNewSubscriberHandler(newSubscriber, lists)

		ctx.JSON(http.StatusAccepted, gin.H{"InsertedId": result})
		ctx.Done()
	}
}
