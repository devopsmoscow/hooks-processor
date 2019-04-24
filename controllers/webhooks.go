package controllers

import (
	"fmt"
	"github.com/Sirupsen/logrus"
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

		logrus.Info(results)

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DialogflowWebhook() {
	http.HandleFunc("/webhook", WebhookRouterHandler)

}
