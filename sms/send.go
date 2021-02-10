package sms

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/michaelwp/go-get-rabbitmq/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Sms struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

func (s Sms) Send() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	var sendMessages model.SendMessages
	var messages model.Messages
	var destinations model.Destinations

	destinations.To = s.To
	messages.From = s.From
	messages.Text = s.Text
	messages.Destinations = append(messages.Destinations, destinations)
	sendMessages.Messages = append(sendMessages.Messages, messages)

	b, err := json.Marshal(sendMessages)
	if err != nil {
		return err
	}

	body := strings.NewReader(string(b))

	req, err := http.NewRequestWithContext(
		ctx, "POST", os.Getenv("INFOBIP_URL"), body)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", os.Getenv("INFOBIP_AUTHORIZATION"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if res.StatusCode != 200 {
		err = errors.New("error sent sms")
		return err
	}

	var sentResponse model.SentResponse

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyByte, &sentResponse)
	if err != nil {
		return err
	}

	return nil
}
