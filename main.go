package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/juhanas/golang-training-advanced/internal/handlers"
)

var appEnv = os.Getenv("appEnv")

func main() {
	setupLogger()

	router := setupRouter()

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	// Init Gin router with default values
	router := gin.Default()

	// Logger middleware logs to os.StdOut
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	initRoutes(router)

	return router
}

func setupLogger() {
	logLevel := slog.LevelInfo
	opts := &slog.HandlerOptions{
		// Include log source (file & line number) if using debug
		AddSource: logLevel == slog.LevelDebug,
		// Set desired logging level
		Level: logLevel,
	}
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if appEnv == "production" {
		// Set logging output to happen as JSON
		// This affects the whole project
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func initRoutes(router *gin.Engine) {
	router.GET("/find-word", handlers.FindWordHandler)
	router.POST("/secrets", handlers.AddSecret)
	router.GET("/secrets/:name", handlers.GetSecret)
	router.GET("/secrets/count", handlers.CountSecrets)
}
