package requests

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"hooks-processor/parser"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Authenticated bool
	Permissions   string
	Action        string
	Payload       parser.Payload
}

func SendMessage(message Message, url string) {
	fmt.Println("URL:>", url)
	json, err := json2.Marshal(message)
	var jsonStr = []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		logrus.Error(err.Error())
	}
	req.Header.Set("X-Custom-Header", "ololosh")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
