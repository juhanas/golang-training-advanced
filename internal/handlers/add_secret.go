package handlers

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/juhanas/golang-training-advanced/pkg/secret"
)

// Internal struct to allow reading data from the post request
type secretDataStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Instantiate a map of secrets.
// Note: If we wanted to store different type of data, we could
// simply replace *secreter.Secret with some other type
// that implements the secreter.Secreter interface and
// everything else in this package would work the same!
var secrets = map[string]*secret.Secreter[string]{}

var counts = map[string]int{
	"created": 0,
	"read":    0,
}

var mutex sync.RWMutex

// Gin route handler for adding a secret to the memory-db.
// The response (success or fail) will be written into the context.
// Expects payload with 'name' and 'value' (string) in the request body.
func AddSecret(c *gin.Context) {
	var secretData secretDataStruct
	if err := c.BindJSON(&secretData); err != nil {
		c.String(http.StatusBadRequest, "unable to parse request body")
		slog.Warn(
			"validation error: AddSecret",
			slog.String("error", err.Error()),
		)
		return
	}

	slog.Debug(
		"incoming request: GetSecret",
		slog.String("name", secretData.Name),
	)

	if secretData.Name == "" || secretData.Value == "" {
		c.String(http.StatusBadRequest, "name, value can't be empty")
		slog.Warn(
			"validation error: AddSecret",
			slog.String("name", secretData.Name),
		)
		return
	}

	if _, exists := secrets[secretData.Name]; exists {
		c.String(http.StatusBadRequest, "secret with provided name already exists")
		slog.Warn(
			"validation error: AddSecret",
			slog.String("name", secretData.Name),
		)
		return
	}

	// By defining the variable's type as secreter.Secreter,
	// we can easily change the implementation of the secret
	var secretItem secret.Secreter[string]
	// This line defines which implementation of the secreter.Secreter
	// interface we want to use. In this case, it's the secreter.String
	secretItem = secret.NewString(secretData.Name)
	// Regardless of the implementation, the Encrypt function will always
	// work the same way and this line does not need to be modified
	encryptedData, err := secretItem.Encrypt(secretData.Value)
	if err != nil {
		c.String(http.StatusInternalServerError, "something went wrong")
		slog.Error(
			"server error from secreter.NewSecret in function: handlers.AddSecret",
			slog.String("error", err.Error()),
		)
		return
	}

	secrets[secretItem.GetName()] = &secretItem
	mutex.Lock()
	counts["created"] = counts["created"] + 1
	mutex.Unlock()

	c.String(http.StatusCreated, encryptedData)
}
