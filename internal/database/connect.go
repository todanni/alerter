package database

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/todanni/alerter/internal/config"
)

func Connect(cfg config.Config) (*amqp.Channel, error) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:5672/", cfg.RMQUser, cfg.RMQPassword, cfg.RMQHost)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}
