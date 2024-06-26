package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindWordHandler(t *testing.T) {
	dirPathOrig := dirPath
	defer func() {
		// reset the dirPath to original to avoid unexpected side effects in other tests
		dirPath = dirPathOrig
	}()

	// dirPath must be set in reference to the location where the test is run - now "./internal/handlers"
	dirPath = "../../data"

	responseText := fmt.Sprintf("Found word '%s' %d times", "cat", 2206)
	runFindWordTest(t, "cat", responseText, 200)
}

func TestFindWordHandlerFileNotFound(t *testing.T) {
	dirPathOrig := dirPath
	defer func() {
		dirPath = dirPathOrig
	}()
	dirPath = "./not-found"

	responseText := "open ./not-found: The system cannot find the file specified."
	runFindWordTest(t, "cat", responseText, 500)
}

func runFindWordTest(t *testing.T, wordToFind, responseText string, statusCode int) {
	// Initialize a mock context with http recorder for gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL: &url.URL{},
	}

	q := req.URL.Query()
	q.Add("word", wordToFind)
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	FindWordHandler(ctx)

	assert.Equal(t, statusCode, w.Code)
	assert.Equal(t, responseText, w.Body.String())
}
