package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	// Import package for its side-effects only (now init)
	_ "github.com/juhanas/golang-training-advanced/pkg/testhelpers"
)

func TestFindWordRouteRecursive(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/find-word?word=cat", nil)
	// mocks the server for a single call
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `Found word 'cat' 2206 times with concurrent false`, w.Body.String())
}

func TestFindWordRouteRecursiveEmptyWord(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/find-word", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `missing query param: word`, w.Body.String())
}

func TestFindWordRouteConcurrent(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/find-word?word=cat&concurrent=true", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `Found word 'cat' 2206 times with concurrent true`, w.Body.String())
}
