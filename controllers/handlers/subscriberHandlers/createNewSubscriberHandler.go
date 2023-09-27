package subscriberHandlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TranHungKT/email_management/constants"
	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"github.com/TranHungKT/email_management/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func sendOptinConfirmationEmail(toEmail string, list []models.List, done chan bool) {
	var nonce, cipherEmail = utils.EncryptCipher(toEmail)
	var startedTime = time.Now().Local().Unix()

	var optinURL = fmt.Sprintf(constants.OptinURLFormat, nonce, cipherEmail, startedTime)

	var templateData = struct {
		Name     string
		Lists    []models.List
		OptinURL string
	}{
		Name:     toEmail,
		Lists:    list,
		OptinURL: optinURL,
	}

	utils.SendEmails([]string{toEmail}, constants.EMAIL_CONFIRMATION_OPTIN_TITLE, constants.EMAIL_CONFIRMATION_OPTIN_TEMPLATE, templateData, done)
}

func getSubscribedAndDoubleOptinLists(lists []models.List) ([]models.SubscribedList, []models.List) {
	var subscribedLists = make([]models.SubscribedList, 0)
	var listsWithDoubleOptin = make([]models.List, 0)
	for _, list := range lists {
		subStatus := models.SubscriptionStatusConfirmed

		if list.Optin == models.ListOptinDouble {
			subStatus = models.SubscriptionStatusUnConfirmed
			listsWithDoubleOptin = append(listsWithDoubleOptin, list)

		}

		subscribedList := models.SubscribedList{
			ListId:             list.Id,
			SubscriptionStatus: subStatus,
		}
		subscribedLists = append(subscribedLists, subscribedList)
	}
	return subscribedLists, listsWithDoubleOptin
}
func timeConsumingTask(done chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("Time-consuming task finished")
	done <- true // Signal that the task is done
}

func CreateNewSubscriberHandler(newSubscriber models.NewSubscriberRequestPayload, lists []models.List, done chan bool) (primitive.ObjectID, error) {
	if newSubscriber.Status == "" {
		newSubscriber.Status = models.SubscriberStatusEnabled
	}
	subscribedLists, listsWithDoubleOptin := getSubscribedAndDoubleOptinLists(lists)

	if len(listsWithDoubleOptin) != 0 {
		go sendOptinConfirmationEmail(newSubscriber.Email, listsWithDoubleOptin, done)
		// go timeConsumingTask(done)
	}

	newSubscriber.Name = strings.TrimSpace(newSubscriber.Name)

	var subscriber = models.Subscriber{
		Email:      newSubscriber.Email,
		Name:       newSubscriber.Name,
		Attributes: newSubscriber.Attributes,
		Status:     newSubscriber.Status,
		Lists:      subscribedLists,
	}

	result, err := database.SubscriberCollection().InsertOne(context.TODO(), &subscriber)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
