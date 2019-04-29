package parser

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
)

var id, username string

type Payload struct {
	Id         int
	Username   string
	Parameters interface{}
}

func ParseWebhook(webhookStrings []string) Payload {
	var messageToBeSent Payload
	for _, webhookString := range webhookStrings {
		var webhook map[string]interface{}
		err := json.Unmarshal([]byte(webhookString), &webhook)
		if err != nil {
			logrus.Error(err.Error())
		}
		for key, value := range webhook {
			if key == "originalRequest" {
				origignalRequest := value.(map[string]interface{})
				for reqKey, reqValue := range origignalRequest {
					if reqKey == "data" {
						data := reqValue.(map[string]interface{})
						for dataKey, dataValue := range data {
							if dataKey == "message" {
								message := dataValue.(map[string]interface{})
								for messageKey, messageValue := range message {
									if messageKey == "from" {
										for fromKey, fromVal := range messageValue.(map[string]interface{}) {
											if fromKey == "id" {
												id := fromVal.(float64)
												messageToBeSent.Id = int(id)
											} else if fromKey == "username" {
												messageToBeSent.Username = fromVal.(string)
											}
										}

									}
								}
							}
						}
					}
				}
			} else if key == "result" {
				result := value.(map[string]interface{})
				for resultKey, resultValue := range result {
					if resultKey == "parameters" {
						messageToBeSent.Parameters = resultValue
					}
				}
			}

		}

	}

	return messageToBeSent
}
