package config

// Constants declarations.
const (
	ProductionScope = "production"
	DevelopScope    = "develop"
	LocalScope      = "local"

	// Environments.
	envHost     string = "DB_HOST_%s"
	envUsername string = "DB_USERNAME_%s"
	envDBName   string = "DB_DBNAME_%s"
	envPassword string = "DB_PASSWORD_%s"

	// SQL constants.
	DBDriver = "mysql"
	DBSNMask = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true"
)

var (
	// EnvVariables var to use load env variables.
	Env EnvVariables

	// DataBase configurations.
	DBConfiguration DBConfig
	// DatSource Name.
	DBDSN string
)
