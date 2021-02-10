package cron

import (
	"github.com/michaelwp/go-get-rabbitmq/rabbitmq"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func ReceiveMsg() (*cron.Cron,  error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	rmqConfig := rabbitmq.RMQConfig{
		Url:   os.Getenv("RABBITMQ_URL"),
		Queue: "hello",
	}

	c := cron.NewWithLocation(loc)
	err = c.AddFunc("* * * * * *", func() {
		rmqConfig.Queue = "otp"
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
