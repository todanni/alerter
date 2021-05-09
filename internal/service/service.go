package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/todanni/alerter/internal/config"
	"github.com/todanni/alerter/pkg/alerter"
)

func NewAlerterService(cfg config.Config, client http.Client) alerter.Service {
	return &alerterService{
		cfg:    cfg,
		client: &client,
	}
}

type alerterService struct {
	client *http.Client
	cfg    config.Config
}

func (s *alerterService) SendLoginAlert() error {
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

func (s *alerterService) SendRegisterAlert() error {
	message := alerter.Message{
		Embed: []*alerter.MessageEmbed{{
			Title:       "New register request received",
			Description: "",
		},
		}}

	return s.sendMessage(message, s.cfg.DiscordRegisterHook)
}

func (s *alerterService) SendActivationAlert() error {
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
	log.Print(string(body))
	return nil
}
