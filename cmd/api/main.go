package main

import (
	"github.com/nicolasburgueno/first_project/internal/app"
	"github.com/nicolasburgueno/first_project/internal/config"
	"github.com/nicolasburgueno/first_project/pkg/database"
)

func main() {
	// Connect to database
	db := database.GetDB()
	// Create application
	app := app.NewApp(db)
	// Run it
	app.Start(config.Env.Port)
}
