package container

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage/engine"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/replication"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal"
	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/log"
	network "github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_server"
)

var configFileName = ".config.yaml"

// AppConfig contains the complete application configuration.
type AppConfig struct {
	Engine      *engine.Config      `yaml:"engine"`
	Network     *network.Config     `yaml:"network"`
	WAL         *wal.Config         `yaml:"wal"`
	Replication *replication.Config `yaml:"replication"`
	Logging     *log.Config         `yaml:"logging"`
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
