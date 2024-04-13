package config

// EnvVariables describes the environment variables expected by this application.
type EnvVariables struct {
	Scope       string
	Port        string
	Application string
}

// DBConfig describes the DB config.
type DBConfig struct {
	Host     string
	Password string
	Username string
	DBName   string
}
