package container

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

var configFileName = ".config.yaml"

// NetworkConfig contains TCP server configuration settings.
type NetworkConfig struct {
	Address            string         `yaml:"address"`
	IdleTimeout        *time.Duration `yaml:"idle_timeout"`
	MaxMessageSize     *string        `yaml:"max_message_size"`
	MaxConnectionsSize *uint          `yaml:"max_connections"`
}

// WALConfig contains Write-Ahead Logging configuration settings.
type WALConfig struct {
	DirectoryPath        string        `yaml:"directory"`
	MaxSegmentSize       string        `yaml:"segment_size"`
	FlushingBatchTimeout time.Duration `yaml:"flushing_batch_timeout"`
	FlushingBatchSize    int           `yaml:"flushing_batch_size"`
}

// AppConfig contains the complete application configuration.
type AppConfig struct {
	Network NetworkConfig `yaml:"network"`
	WAL     WALConfig     `yaml:"wal"`
}

// NewConfig loads and parses the configuration from .config.yaml file.
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
