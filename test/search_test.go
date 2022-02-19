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

func TestMain(m *testing.M) {
	config()
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}
func TestPingRouter(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func TestSimpleSearch(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュー")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"","names":["カイリュー"]}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestPartialSearch(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュー")
	q.Add("hits", "..x..")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"","names":["リザードン","バタフリー","オニドリル","ニドリーナ","ニドリーノ","ダグトリオ","オコリザル","ワンリキー","ゴーリキー","カイリキー","ドードリオ","スリーパー","ビリリダマ","ベロリンガ","バリヤード","フリーザー","ミニリュウ","ハクリュー","カイリュー"]}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHitsSearch(t *testing.T) {
	_, router := con.CreateConnection()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/poke", nil)
	q := req.URL.Query()
	q.Add("name", "カイリュー")
	q.Add("hits", "o.x..")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := `{"response":{"error":"","names":["カイリキー","カイリュー"]}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func config() {
	apath, _ := filepath.Abs("../")
	err := os.Chdir(apath)
	if err != nil {
		log.Fatal(err)
	}
}
