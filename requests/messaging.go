package requests

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"hooks-processor/parser"
	"net/http"
)

type Message struct {
	Authenticated bool
	Permissions   string
	Action        string
	Payload       parser.Payload
}

type Services struct {
	List []struct {
		URL     string   `mapstructure:"url"`
		Service string   `mapstructure:"service"`
		Actions []string `mapstructure:"actions"`
	} `mapstructure:"list"`
}

func SendMessage(message Message) {
	url, err := getUrl(message.Action)
	json, err := json2.Marshal(message)
	var jsonStr = []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		logrus.Error(err.Error())
	} else {
		req.Header.Set("X-Custom-Header", "ololosh")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()
		}
	}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func getUrl(action string) (string, error) {
	var url string
	for key, value := range viper.Get("services").(map[string]interface{}) {
		if key == "list" {
			for _, service := range value.([]interface{}) {
				var matchFlag bool
				var tempUrl string
				for servKey, servValue := range service.(map[interface{}]interface{}) {
					if servKey == "actions" {
						for _, actionFromConfig := range servValue.([]interface{}) {
							if actionFromConfig == action {
								matchFlag = true
							}
						}
					}
					if servKey == "url" {
						tempUrl = fmt.Sprintf("%v", servValue)
					}
				}
				if matchFlag == true {
					url = tempUrl
				}
			}
		}
	}
	if url == "" {
		logrus.Error("Action or URL not found")
		return "", &errorString{"Action or URL not found"}
	}
	logrus.Info("Integration URL is ", url)
	return url, nil
}
