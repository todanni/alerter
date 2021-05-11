package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromEnv(t *testing.T) {
	config, err := NewFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "localhost", config.RMQHost)
	assert.Equal(t, "docker", config.RMQUser)
	assert.Equal(t, "docker", config.RMQPassword)
}
