package cron

import (
	"github.com/michaelwp/go-get-rabbitmq/rabbitmq"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"time"
)

func ReceiveMsg() (*cron.Cron,  error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	c := cron.NewWithLocation(loc)

	rmqConfig := rabbitmq.RMQConfig{
		Url:   "amqp://guest:guest@localhost:5672/",
		Queue: "hello",
	}

	err = c.AddFunc("* * * * * *", func() {
		err := rmqConfig.Receive()
		if err != nil {
			logrus.Error(err)
		}
	})

	if err != nil {
		return nil, err
	}
	c.Start()

	return c, nil
}
