// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Hello World", resp.Body.String())
}

func TestGetGreet(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/greet", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// You can add more assertions based on the expected content of the greeting page.
}

func TestGetAllMovies(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/movie", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// You can add more assertions based on the expected content of the all-movies page.
}

// Add more tests as needed

func setupRouter() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", getHello)
	router.GET("/greet", getGreet)
	router.GET("/movie", getAllMovies)
	return router
}
