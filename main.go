package main

import (
	"github.com/gin-gonic/gin"

	"github.com/juhanas/golang-training-advanced/internal/handlers"
)

func main() {
	router := setupRouter()

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	// Init Gin router with default values
	router := gin.Default()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	initRoutes(router)

	return router
}

func initRoutes(router *gin.Engine) {
	router.GET("/find-word", handlers.FindWordHandler)
}
