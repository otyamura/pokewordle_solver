package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	con "github.com/otyamura/pokewordle_solver/internal/connection"
	"github.com/stretchr/testify/assert"
)

func TestPingRouter(t *testing.T) {
	config()
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// ...
	assert.Equal(t, w.Body.String(), "{\"message\":\"pong\"}")
}

func TestSimpleSearch(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュー")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"names":["カイリュー"]}`
	assert.Equal(t, 200, w.Code)
	// ...
	assert.Equal(t, w.Body.String(), expected)
}

func config() {
	apath, _ := filepath.Abs("../")
	err := os.Chdir(apath)
	if err != nil {
		log.Fatal(err)
	}
}
