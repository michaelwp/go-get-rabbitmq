package main

import (
	"github.com/michaelwp/go-get-rabbitmq/cron"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main()  {
	logrus.Println(" [*] Waiting for messages. To exit press CTRL+C")

	c, err := cron.ReceiveMsg()
	go c.Start()

	if err != nil {
		logrus.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
