package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/secret"
)

func TestIntegrations(t *testing.T) {
	callsToGet := 5000

	defer func() {
		secrets = map[string]*secret.Secreter{}
		counts["read"] = 0
		counts["created"] = 0
	}()

	secretName := "testSecret"
	secretValue := "abc"
	secretItem := secret.NewString(secretName)
	_, err := secretItem.Encrypt(secretValue)
	if err != nil {
		panic(err)
	}

	var actualSecret secret.Secreter
	actualSecret = secretItem
	secrets[secretName] = &actualSecret

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, eng := gin.CreateTestContext(w)

	eng.GET("/secrets/:name", GetSecret)

	req, err := http.NewRequest("GET", "/secrets/"+secretName, nil)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for range callsToGet {
		wg.Add(1)
		go func() {
			eng.ServeHTTP(w, req)
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, secretValue, w.Body.String()[0:3])
	assert.Equal(t, callsToGet, counts["read"])
}
