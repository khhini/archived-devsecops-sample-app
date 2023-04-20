package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHello(t *testing.T) {
	router := SetupRouter()
	expectedRes := `{"message":"Hello, World!."}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedRes, w.Body.String())
}

func TestHelloByName(t *testing.T) {
	router := SetupRouter()
	expectedRes := `{"message":"Hello, kiki!."}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/kiki", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedRes, w.Body.String())
}
