package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin route handler for retrieving a secret from the memory-db.
// The response will be written into the context.
// Expects path param with 'name' (string) in the request.
func GetSecret(c *gin.Context) {
	name := c.Param("name")
	slog.Info(
		"incoming request: GetSecret",
		slog.String("name", name),
		slog.Any("URL", c.Request.URL),
	)

	secret, exists := secrets[name]
	if !exists {
		c.String(http.StatusNotFound, "secret with provided name does not exists")
		slog.Warn(
			"not found error: GetSecret",
			slog.String("name", name),
		)
		return
	}

	data, err := (*secret).Decrypt()
	if err != nil {
		c.String(http.StatusInternalServerError, "secret could not be obtained")
		slog.Warn(
			"error: secret.Decrypt",
			slog.String("error", err.Error()),
		)
		return
	}

	counts["read"] = counts["read"] + 1

	c.String(http.StatusOK, data)
}
