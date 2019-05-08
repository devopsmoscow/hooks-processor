package parser

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
)

func ParseWebhook(webhookStrings []string) Payload {
	logrus.Info("Payload parser initiated")
	var messageToBeSent Payload
	for _, webhookString := range webhookStrings {
		var webhook TelegramWebhoook
		err := json.Unmarshal([]byte(webhookString), &webhook)
		if err != nil {
			logrus.Error(err.Error())
		}
		from := webhook.OriginalDetectIntentRequest.Payload.Data.Message.From
		messageToBeSent.Username = from.Username
		logrus.Info(messageToBeSent.Username)
		messageToBeSent.Id = int(from.ID)
		logrus.Info(messageToBeSent.Id)
		messageToBeSent.Parameters = webhook.QueryResult.Parameters
		logrus.Info(messageToBeSent.Parameters)
		logrus.Info("Parsing finished")

	}

	return messageToBeSent
}
