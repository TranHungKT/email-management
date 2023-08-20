package subscriberControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfirmOptionPayload struct {
	Email string `json:"email" validate:"email,required"`
}

func ConfirmOptinController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var confirmOptionPayload ConfirmOptionPayload
		err := utils.BindJSONAndValidateByStruct(ctx, &confirmOptionPayload)
		if err != nil {
			return
		}

		var updatedDocument bson.M

		arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"list.subscriptionStatus": "unconfirmed"}}}

		var returnDocument options.ReturnDocument = 1
		filterOption := options.FindOneAndUpdateOptions{
			ArrayFilters:   &arrayFilters,
			ReturnDocument: &returnDocument,
		}

		err = database.SubscriberCollection().FindOneAndUpdate(
			context.TODO(),
			bson.D{bson.E{Key: "email", Value: confirmOptionPayload.Email}},
			bson.M{"$set": bson.M{"lists.$[list].subscriptionStatus": models.SubscriptionStatusConfirmed}}, &filterOption).Decode(&updatedDocument)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, updatedDocument)

	}
}