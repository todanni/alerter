package config

import "github.com/caarlos0/env/v6"

// Config contains the env variables needed to run the servers
type Config struct {
	RMQUser     string `env:"RMQ_USER,required"`
	RMQPassword string `env:"RMQ_PASSWORD,required"`
	RMQHost     string `env:"RMQ_HOST,required"`

	DiscordRegisterHook   string `env:"DISCORD_REGISTER,required"`
	DiscordActivationHook string `env:"DISCORD_VERIFY,required"`
	DiscordLoginHook      string `env:"DISCORD_LOGIN,required"`
}

func NewFromEnv() (Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
