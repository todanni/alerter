package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerter/internal/config"
	"github.com/todanni/alerter/internal/queue"
	"github.com/todanni/alerter/internal/service"
)

func main() {
	// Read config
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	ch, err := queue.Connect(cfg)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"alerts",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	s := service.NewAlerterService(cfg, http.Client{})

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			switch string(d.Body) {
			case "login":
				err = s.SendLoginAlert()
			case "register":
				err = s.SendRegisterAlert()
			case "verified":
				err = s.SendActivationAlert()
			}
			if err != nil {
				log.Error(err)
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
