package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"hooks-processor/parser"
	"hooks-processor/requests"
	"io/ioutil"
	"net/http"
)

type Response struct {
	fulfillmentText string
}

func WebhookRouterHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Received a request")
	var results []string
	if r.Method == "POST" {
		logrus.Info("Request method is POST")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))
		payload := parser.ParseWebhook(results)
		fmt.Println(payload)
		var message requests.Message
		message.Payload = payload
		message.Authenticated = false
		message.Permissions = "admin"
		logrus.Info(message.Payload.Parameters)
		logrus.Info(message.Payload.Username)
		for paramKey, paramValue := range payload.Parameters.(map[string]interface{}) {
			if paramKey == "action" {
				logrus.Info("Found 'action' key")
				message.Action = paramValue.(string)
			}
		}
		logrus.Info("Requesting integration")
		requests.SendMessage(message)
		var response Response
		response.fulfillmentText = "Test Response"
		js, err := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		_, wErr := w.Write(js)
		if wErr != nil {
			logrus.Error(wErr.Error())
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DialogflowWebhook() {
	http.HandleFunc("/webhook", WebhookRouterHandler)

}
