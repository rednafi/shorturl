package tests

import (
	"encoding/json"
	"fmt"
	"github.com/rednafi/shorturl/src"
	"testing"
)

func TestRequest(t *testing.T) {
	url := "http://www.example.com/long/url"
	req := src.Request{Url: url}

	// Check if the request url matches
	if req.Url != url {
		t.Error("Request Url does not match")
	}

}

func TestRequestJson(t *testing.T) {
	url := "http://www.example.com/long/url"
	req := src.Request{Url: url}

	// Check if encoded json matches
	want := fmt.Sprintf(`{"url":"%s"}`, url)
	j, _ := json.Marshal(req)
	got := string(j)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestResponse(t *testing.T) {
	url := "http://www.example.com/long/url"
	res := src.Response{TinyUrl: url}
	if res.TinyUrl != url {
		t.Error("Response TinyUrl does not match")
	}
}

func TestResponseJson(t *testing.T) {
	url := "http://www.example.com/long/url"
	res := src.Response{TinyUrl: url}

	// Check if encoded json matches
	want := fmt.Sprintf(`{"tinyurl":"%s"}`, url)
	j, _ := json.Marshal(res)
	got := string(j)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
