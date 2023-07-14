package subscriberHandlers

import (
	"context"
	"strings"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewSubscriberHandler(newSubscriber models.NewSubscriberRequestPayload, lists []models.List) (interface{}, error) {
	if newSubscriber.Status == "" {
		newSubscriber.Status = models.SubscriberStatusEnabled
	}

	var subscribedLists = make([]models.SubscribedList, 0)

	for _, list := range lists {
		subStatus := models.SubscriptionStatusConfirmed

		if list.Optin == models.ListOptinDouble {
			subStatus = models.SubscriptionStatusUnConfirmed
		}

		subscribedList := models.SubscribedList{
			ListId:             list.Id,
			SubscriptionStatus: subStatus,
		}
		subscribedLists = append(subscribedLists, subscribedList)
	}

	newSubscriber.Name = strings.TrimSpace(newSubscriber.Name)

	result, err := database.SubscriberCollection().InsertOne(context.TODO(), primitive.D{
		primitive.E{Key: "email", Value: newSubscriber.Email},
		primitive.E{Key: "name", Value: newSubscriber.Name},
		primitive.E{Key: "attributes", Value: newSubscriber.Attributes},
		primitive.E{Key: "status", Value: newSubscriber.Status},
		primitive.E{Key: "lists", Value: subscribedLists},
	})

	if err != nil {
		return result.InsertedID, err
	}

	return result.InsertedID, nil
}
