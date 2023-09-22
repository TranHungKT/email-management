package subscriberControllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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

		nonce := ctx.Param("nonceKey")
		cipherEmail := ctx.Param("cipherEmailKey")
		startedTime := ctx.Param("startedTime")

		sTime, err := strconv.ParseInt(startedTime, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(sTime, 0)
		maximumTime := tm.Add(time.Hour * 3)

		if maximumTime.Before(time.Now()) {
			fmt.Print("OK it is invalid")
			ctx.HTML(http.StatusOK, "outOfDateURL.html", gin.H{})

			return
		}

		email := utils.DecryptCipher(nonce, cipherEmail)

		var updatedDocument bson.M

		arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"list.subscriptionStatus": "unconfirmed"}}}

		var returnDocument options.ReturnDocument = 1
		filterOption := options.FindOneAndUpdateOptions{
			ArrayFilters:   &arrayFilters,
			ReturnDocument: &returnDocument,
		}

		err = database.SubscriberCollection().FindOneAndUpdate(
			context.TODO(),
			bson.D{bson.E{Key: "email", Value: email}},
			bson.M{"$set": bson.M{"lists.$[list].subscriptionStatus": models.SubscriptionStatusConfirmed}}, &filterOption).Decode(&updatedDocument)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.HTML(http.StatusOK, "successSubscription.html", gin.H{})
	}
}
