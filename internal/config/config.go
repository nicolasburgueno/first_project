package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// init init function.
func init() {
	Load()
}

// Load Loads the configurations.
func Load() {

	// Load Environments
	LoadEnv()

	switch scope := Env.Scope; {
	case strings.HasPrefix(scope, ProductionScope):
		loadProductionValues()
	case strings.HasPrefix(scope, DevelopScope):
		loadDevelopValues()
	default:
		loadLocalValues()
	}

	// DB
	PrepareDBConn()
}

// LoadEnv load environments values.
func LoadEnv() {
	// Process to load environments.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Env.Application = os.Getenv("APPLICATION")
	Env.Scope = os.Getenv("SCOPE")
	Env.Port = os.Getenv("PORT")
}

// loadProductionValues load configurations for production environment.
func loadProductionValues() {
	DBConfiguration = DBConfig{
		Host:     fmt.Sprintf(envHost, "PROD"),
		DBName:   fmt.Sprintf(envDBName, "PROD"),
		Username: fmt.Sprintf(envUsername, "PROD"),
		Password: fmt.Sprintf(envPassword, "PROD"),
	}
}

// loadDevelopValues load configurations for develop environment.
func loadDevelopValues() {
	DBConfiguration = DBConfig{
		Host:     fmt.Sprintf(envHost, "DEV"),
		DBName:   fmt.Sprintf(envDBName, "DEV"),
		Username: fmt.Sprintf(envUsername, "DEV"),
		Password: fmt.Sprintf(envPassword, "DEV"),
	}
}

// loadLocalValues load configurations for local environment.
func loadLocalValues() {
	DBConfiguration = DBConfig{
		Host:     fmt.Sprintf(envHost, "LOCAL"),
		DBName:   fmt.Sprintf(envDBName, "LOCAL"),
		Username: fmt.Sprintf(envUsername, "LOCAL"),
		Password: fmt.Sprintf(envPassword, "LOCAL"),
	}
}

// PrepareDBConn prepare connection to the database.
func PrepareDBConn() {
	envKeyHost := os.Getenv(DBConfiguration.Host)
	envKeyDBName := os.Getenv(DBConfiguration.DBName)
	envKeyUsername := os.Getenv(DBConfiguration.Username)
	envKeyPassword := os.Getenv(DBConfiguration.Password)

	DBDSN = fmt.Sprintf(DBSNMask, envKeyUsername, envKeyPassword, envKeyHost, envKeyDBName)
}
