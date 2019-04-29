package controllers

import (
	"fmt"
	"hooks-processor/parser"
	"hooks-processor/requests"
	"io/ioutil"
	"net/http"
)

func WebhookRouterHandler(w http.ResponseWriter, r *http.Request) {
	var results []string
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))

		payload := parser.ParseWebhook(results)
		var message requests.Message
		message.Payload = payload
		message.Authenticated = false
		message.Permissions = "admin"

		for paramKey, paramValue := range payload.Parameters.(map[string]interface{}) {
			if paramKey == "action" {
				message.Action = paramValue.(string)
			}
		}
		var url string
		switch message.Action {
		case "new_meetup":
			url = "https://webhook.site/65c30cd6-674f-4390-a20f-7ed5ea4961b2"
		}
		requests.SendMessage(message, url)
		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DialogflowWebhook() {
	http.HandleFunc("/webhook", WebhookRouterHandler)

}
