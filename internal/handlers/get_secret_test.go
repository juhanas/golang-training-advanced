package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/secret"
)

func TestGetSecret(t *testing.T) {
	defer func() {
		secrets = map[string]*secret.Secreter[string]{}
		counts["read"] = 0
	}()

	secretName := "testSecret"
	secretValue := "abc"
	secretItem := secret.NewString(secretName)
	_, err := secretItem.Encrypt(secretValue)
	if err != nil {
		panic(err)
	}

	// Note: The go static checker might complain about this line
	// but by using this type casting we can make sure that we can
	// use every implementation of the Secreter interface without
	// having to change this part of the code.
	var actualSecret secret.Secreter[string]
	actualSecret = secretItem // Type cast to the interface
	secrets[secretName] = &actualSecret

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
