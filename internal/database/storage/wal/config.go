package wal

import "time"

// Config contains Write-Ahead Logging configuration settings.
type Config struct {
	DirectoryPath        string        `yaml:"directory"`
	MaxSegmentSize       string        `yaml:"segment_size"`
	FlushingBatchTimeout time.Duration `yaml:"flushing_batch_timeout"`
	FlushingBatchSize    int           `yaml:"flushing_batch_size"`
}
