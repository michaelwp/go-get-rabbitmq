package rabbitmq

import (
	"encoding/json"
	sms2 "github.com/michaelwp/go-get-rabbitmq/sms"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
)

type RMQConfig struct {
	Url   string `json:"url"`
	Queue string `json:"queue"`
}

func (r RMQConfig) Receive() error {
	conn, err := amqp.Dial(r.Url)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		r.Queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var sms sms2.Sms

			err = json.Unmarshal(d.Body, &sms)
			if err != nil {
				logrus.Error(err)
			}

			logrus.Info(sms)

			err = sms.Send()
			if err != nil {
				logrus.Error(err)
			}

			log.Printf("Received a message: %s", d.Body)
		}
	}()

	return nil
}
