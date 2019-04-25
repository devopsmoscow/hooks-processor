package parser

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
)

var id, username string

func ParseWebhook(webhookStrings []string) string {
	for _, webhookString := range webhookStrings {
		var webhook map[string]interface{}
		err := json.Unmarshal([]byte(webhookString), &webhook)
		if err != nil {
			logrus.Error(err.Error())
		}
		fmt.Println(webhook)
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
												fmt.Println(fromVal)
											}
										}

									}
								}
							}
						}
					}
				}
			}
		}

	}

	//browsersParsed := result["browsersParsed"].(map[string]interface{})
	//var browsers []Browser

	return ""
}
