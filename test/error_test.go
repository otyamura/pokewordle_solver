package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	con "github.com/otyamura/pokewordle_solver/internal/connection"
	"github.com/stretchr/testify/assert"
)

func TestGenValidation(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("less", "0")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"invalid generation","names":null}}`
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHitsLength(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュー")
	q.Add("hits", "..x.")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"poke name or hits not match len","names":null}}`
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestNameLength(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュ")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"poke name or hits not match len","names":null}}`
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestNameValidation(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"poke name or hits not match len","names":null}}`
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
