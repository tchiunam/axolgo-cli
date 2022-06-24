package types

// Structure of logging configuration
type AxolgoConfigLogging struct {
	// Log level verbosity
	LogLevelVerbosity int `mapstructure:"log-level-verbosity"`
}

// Structure of AWS configuration
type AxolgoConfigAWS struct {
	// AWS region
	Region string `mapstructure:"region"`
}

// Structure of GCP configuration
type AxolgoConfigGCP struct {
	// Google application credentials file
	GoogleApplicationCredentials string `mapstructure:"google-application-credentials"`
	Zone                         string `mapstructure:"zone"`
}

// Structure of axolgo configuration
type AxolgoConfig struct {
	// logging configuration
	Logging AxolgoConfigLogging `mapstructure:"logging"`
	AWS     AxolgoConfigAWS     `mapstructure:"aws"`
	GCP     AxolgoConfigGCP     `mapstructure:"gcp"`
}
