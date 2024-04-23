package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/secreter"
)

func TestAddSecret(t *testing.T) {
	defer func() {
		secrets = map[string]*secreter.Secret{}
		counts["created"] = 0
	}()

	secretName := "testSecret"
	secretValue := "abcd"

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	newSecret := []byte(`{"name": "testSecret", "value": "abcd"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(newSecret))
	if err != nil {
		panic(err)
	}

	ctx.Request = req
	AddSecret(ctx)

	assert.Equal(t, 201, w.Code)
	assert.True(t, secretValue != w.Body.String(), fmt.Sprintf("Strings should be different: %s - %s", secretValue, w.Body.String()))

	val, ok := secrets[secretName]
	assert.Equal(t, true, ok)
	assert.Equal(t, counts["created"], 1)

	if ok {
		assert.NotSame(t, secretValue, val)
		// Verify the data was encrypted correctly
		data, err := secrets[secretName].Decrypt()
		assert.Nil(t, err)
		assert.Equal(t, secretValue, string(data))
	}
}
