package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCountSecrets(t *testing.T) {
	counts["read"] = 1
	counts["created"] = 3
	defer func() {
		counts["read"] = 0
		counts["created"] = 0
	}()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, eng := gin.CreateTestContext(w)

	eng.GET("/secrets/count", CountSecrets)

	req, err := http.NewRequest("GET", "/secrets/count", nil)
	if err != nil {
		panic(err)
	}

	eng.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data CountData
	err = json.Unmarshal(w.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.Equal(t, 1, data.Reads)
	assert.Equal(t, 3, data.Writes)
}
