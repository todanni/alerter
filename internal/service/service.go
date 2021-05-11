package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/streadway/amqp"
	"github.com/todanni/alerter/internal/config"
	"github.com/todanni/alerter/pkg/alerter"
	"github.com/todanni/alerts"
)

func NewAlerterService(cfg config.Config, client http.Client, ch *amqp.Channel) alerter.Service {
	return &alerterService{
		cfg:     cfg,
		client:  &client,
		channel: ch,
	}
}

type alerterService struct {
	client  *http.Client
	channel *amqp.Channel
	cfg     config.Config
}

func (s *alerterService) Run() error {
	q, err := s.channel.QueueDeclare("alerts",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}

	msgs, err := s.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}

	for m := range msgs {
		err = s.resolveAlert(m.Body)
		if err != nil {
			log.Print(err)
		}
	}
	return nil
}

func (s *alerterService) resolveAlert(alertBytes []byte) error {
	var al alerts.Alert
	err := json.Unmarshal(alertBytes, &al)
	if err != nil {
		return err
	}

	switch al.Type {
	case "login":
		return s.sendLoginAlert()
	case "register":
		return s.sendRegisterAlert()
	case "verify":
		return s.sendActivationAlert()
	default:
		return errors.New("unknown alert type")
	}
}

func (s *alerterService) sendLoginAlert() error {
	// TODO: get public details for user
	message := alerter.Message{
		Embed: []*alerter.MessageEmbed{{
			Title:       "A user has just logged in",
			Description: "",
			Image: &alerter.MessageEmbedImage{
				URL: "https://i.imgur.com/ibL476q.png",
			},
		},
		}}

	return s.sendMessage(message, s.cfg.DiscordLoginHook)
}

func (s *alerterService) sendRegisterAlert() error {
	message := alerter.Message{
		Embed: []*alerter.MessageEmbed{{
			Title:       "New register request received",
			Description: "",
		},
		}}

	return s.sendMessage(message, s.cfg.DiscordRegisterHook)
}

func (s *alerterService) sendActivationAlert() error {
	// TODO: get public details for user
	message := alerter.Message{
		Embed: []*alerter.MessageEmbed{{
			Title:       " just verified their email!",
			Description: "",
		},
		}}
	return s.sendMessage(message, s.cfg.DiscordActivationHook)
}

func (s *alerterService) sendMessage(message alerter.Message, hook string) error {
	postBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, hook, b)
	req.Header.Add("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	_ = body
	//log.Print(string(body))
	return nil
}
