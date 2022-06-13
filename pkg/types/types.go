package types

// Structure of logging configuration
type AxolgoConfigLogging struct {
	// Log level verbosity
	LogLevelVerbosity int `mapstructure:"log_level_verbosity"`
}

// Structure of AWS configuration
type AxolgoConfigAWS struct {
	// AWS region
	Region string `mapstructure:"region"`
}

// Structure of axolgo configuration
type AxolgoConfig struct {
	// logging configuration
	Logging AxolgoConfigLogging `mapstructure:"logging"`
	AWS     AxolgoConfigAWS     `mapstructure:"aws"`
}
