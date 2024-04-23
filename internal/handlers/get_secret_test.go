package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/secreter"
)

func TestGetSecret(t *testing.T) {
	defer func() {
		secrets = map[string]*secreter.Secret{}
		counts["read"] = 0
	}()

	secretName := "testSecret"
	secretValue := "abc"
	secret := secreter.NewSecret(secretName)
	_, err := secret.Encrypt(secretValue)
	if err != nil {
		panic(err)
	}

	secrets[secretName] = secret

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, eng := gin.CreateTestContext(w)

	eng.GET("/secrets/:name", GetSecret)

	req, err := http.NewRequest("GET", "/secrets/"+secretName, nil)
	if err != nil {
		panic(err)
	}

	eng.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, secretValue, w.Body.String())
	assert.Equal(t, 1, counts["read"])
}
