package in_memory

type Config struct {
	PartitionsCount int `envconfig:"IN_MEMORY_ENGINE_PARTITIONS_COUNT" required:"true"`
}