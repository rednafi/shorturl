package tests

import (
	"bytes"
	"encoding/json"
	"github.com/rednafi/shorturl/src"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// Run the test functions
	exitCode := m.Run()

	// Remove the test database
	os.Remove("urls.db")

	os.Exit(exitCode)
}

func TestShortenUrl(t *testing.T) {

	reqBody := `{"url": "https://rednafi.com"}`
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBufferString(reqBody))

	rr := httptest.NewRecorder()

	src.ShortenUrl(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK,
		)
	}

	var resp map[string]string
	json.Unmarshal(rr.Body.Bytes(), &resp)

	// Check response has tinyurl key
	if _, ok := resp["tinyurl"]; !ok {
		t.Errorf("handler returned unexpected body: got %v want %v", resp, "tinyurl")
	}
}
