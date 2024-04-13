package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Router *gin.Engine
}

// NewApp create a new instance of application.
func NewApp(db *sqlx.DB) *App {
	app := &App{
		Router: gin.Default(),
	}

	app.setuRoutes()

	return app
}

// Start the application by configuration and running the Gin-Gonic router.
func (app *App) Start(port string) {
	log.Println("Starting application - Port:", port)
	if err := app.Router.Run(":" + port); err != nil {
		panic("Error running server")
	}
}

// SetupRoutes setup the application routes.
func (app *App) setuRoutes() {
	// Implement me.

	// TODO: Use middlewares.
	app.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
