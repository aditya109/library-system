package models

// Config contains all the configuration-related to server and application.
type Config struct {
	ServerConfig      ServerConfig      `json:"server_config,omitempty"`      // ServerConfig contains server-related configuration.
	ApplicationConfig ApplicationConfig `json:"application_config,omitempty"` // ApplicationConfig contains application-related configuration.
}

// ServerConfig contains server-related configuration.
type ServerConfig struct {
	EnvironmentType string `json:"environment_type,omitempty"` // EnvironmentType indicates the environment being used.
	DevEnvs         Envs   `json:"dev_env,omitempty"`          // DevEnvs contains all the environment variables which is used in the dev setup.
	StagEnvs        Envs   `json:"stag_env,omitempty"`         // StagEnvs contains all the environment variables which is used in the dev setup.
	ProdEnvs        Envs   `json:"prod_envs,omitempty"`        // ProdEnvs contains all the environment variables which is used in the production setup.
}

// Envs contains all the environment variables which is used in the dev setup.
type Envs struct {
	ServerUri    string `json:"server_uri,omitempty"`     // ServerUri defines the uri at which server is initiated.
	ServerPort   string `json:"server_port,omitempty"`    // ServerPort defines the port at which server is initiated.
	WriteTimeout int    `json:"write_timeout,omitempty"`  // WriteTimeout defines the write timeout for the server.
	ReadTimeout  int    `json:"read_timeout,omitempty"`   // ReadTimeout defines the read timeout for the server.
	IsTLSEnabled bool   `json:"is_tls_enabled,omitempty"` // IsTLSEnabled specifies if TLS is to be enabled.
	DBUri        string `json:"db_uri,omitempty"`         // DBUri defines the uri for the database.
}

// ApplicationConfig contains application-related configuration.
type ApplicationConfig struct {
	APIPrefix    string       `json:"api_prefix,omitempty"`    // APIPrefix defines the prefix for the API endpoints
	LevelledLogs LevelledLogs `json:"levelled_logs,omitempty"` // LevelledLogs contains all the application logging-related configuration.
}

// LevelledLogs contains all the application logging-related configuration.
type LevelledLogs struct {
	PersistenceLocation   PersistenceLocation `json:"persistence_location,omitempty"`
	EnableLoggingToFile   bool                `json:"enable_logging_to_file,omitempty"`
	EnableLoggingToStdout bool                `json:"enable_logging_to_stdout,omitempty"`
	EnableColors          bool                `json:"enable_colors,omitempty"`
	EnableFullTimestamp   bool                `json:"enable_full_timestamp,omitempty"`
	OutputFormatter       string              `json:"output_formatter"`
}

// PersistenceLocation contains file-system-related configurations.
type PersistenceLocation struct {
	ContainerDirectory  string   `json:"container_directory,omitempty"`
	TargetFileName      []string `json:"targetfile_name,omitempty"`
	TargetFileExtension string   `json:"targetfile_extension,omitempty"`
}
