package container

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

var configFileName = ".config.yaml"

type NetworkConfig struct {
	Address            string         `yaml:"address"`
	IdleTimeout        *time.Duration `yaml:"idle_timeout"`
	MaxMessageSize     *string        `yaml:"max_message_size"`
	MaxConnectionsSize *uint          `yaml:"max_connections"`
}

type AppConfig struct {
	Network NetworkConfig `yaml:"network"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() (*AppConfig, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(rootDir, configFileName)

	config := &AppConfig{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
