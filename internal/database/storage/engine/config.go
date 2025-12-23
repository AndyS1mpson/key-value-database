package engine

// Config contains storage engine configuration settings.
type Config struct {
	Type             string `yaml:"type"`
	PartitionsNumber *uint  `yaml:"partitions_number"`
}
