package replication

import "time"

// Config contains replication configuration settings.
type Config struct {
	ReplicaType       string        `yaml:"replica_type"`
	MasterAddress     string        `yaml:"master_address"`
	SyncInterval      time.Duration `yaml:"sync_interval"`
	MaxReplicasNumber int           `yaml:"max_replicas_number"`
}
