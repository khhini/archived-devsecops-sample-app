package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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
	req, _ := http.NewRequest("GET", "/hello/kiki", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedRes, w.Body.String())
}

func TestPing(t *testing.T) {
	router := SetupRouter()
	expectedRes := `{"ping":"pong"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedRes, w.Body.String())
}

func TestHealthCheck(t *testing.T) {
	router := SetupRouter()
	os.Setenv("DB_PASSWORD_SECRET_MANAGER_REF", "projects/1078293032607/secrets/sample-app-stage-db-username/versions/latest")
	expectedRes := `{"secrets":"superduperuser"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health_check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedRes, w.Body.String())
}
