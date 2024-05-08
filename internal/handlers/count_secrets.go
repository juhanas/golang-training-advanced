package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountData struct {
	Reads  int
	Writes int
}

// Gin route handler for getting access counts for the secret.
// The response will be written into the context.
func CountSecrets(c *gin.Context) {
	slog.Info(
		"incoming request: CountSecrets",
	)

	// Locking the mutex for reading. This only blocks
	// if there is a write operation at the same time.
	mutex.RLock()
	cd := CountData{
		Reads:  counts["read"],
		Writes: counts["created"],
	}
	mutex.RUnlock()

	c.JSON(http.StatusOK, cd)
}
