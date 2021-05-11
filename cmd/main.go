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
	svc := service.NewAlerterService(cfg, http.Client{}, ch)
	log.Fatal(svc.Run())
}
