package tcp_server

import "time"

// Config contains TCP server configuration settings.
type Config struct {
	Address            string         `yaml:"address"`
	IdleTimeout        *time.Duration `yaml:"idle_timeout"`
	MaxMessageSize     *string        `yaml:"max_message_size"`
	MaxConnectionsSize *uint          `yaml:"max_connections"`
}
