package tests

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/rednafi/shorturl/src"
	"testing"
	"net/http"
	"crypto/tls"
)

func TestValidateUrl(t *testing.T) {
	validUrl := "http://example.com"
	invalidUrl := "totally invalid url"

	if !src.ValidateUrl(validUrl) {
		t.Error("ValidateUrl failed on valid url")
	}

	if src.ValidateUrl(invalidUrl) {
		t.Error("ValidateUrl passed on invalid url")
	}
}

func TestTrimSlash(t *testing.T) {
	url := "http://example.com/"
	expected := "http://example.com"

	result := src.TrimSlash(url)
	if result != expected {
		t.Errorf("TrimSlash failed, expected: %s, got: %s", expected, result)
	}
}

func TestGenerateId(t *testing.T) {
	id := src.GenerateId(10)
	if len(id) != 16 {
		t.Errorf("GenerateId returned id of incorrect length, expected: 10, got: %d", len(id))
	}
}

func TestGenerateHash(t *testing.T) {
	url := "http://example.com"
	// generate expected hash
	hasher := sha1.New()
	hasher.Write([]byte(url))
	expected := hex.EncodeToString(hasher.Sum(nil))

	result := src.GenerateHash(url)
	if result != expected {
		t.Errorf("GenerateHash failed, expected: %s, got: %s", expected, result)
	}
}

func TestGetQualifiedTinyUrl(t *testing.T) {
    t.Run("returns https url if request is tls", func(t *testing.T) {
        req := &http.Request{
            TLS: &tls.ConnectionState{},
            Host: "example.com",
        }
        id := "abc123"

        url := src.GetQualifiedTinyUrl(req, id)

        want := "https://example.com/r/abc123"
        if url != want {
            t.Errorf("got %q, want %q", url, want)
        }
    })

    t.Run("returns http url if request is not tls", func(t *testing.T) {
        req := &http.Request{
            Host: "example.com",
        }
        id := "abc123"

        url := src.GetQualifiedTinyUrl(req, id)

        want := "http://example.com/r/abc123"
        if url != want {
            t.Errorf("got %q, want %q", url, want)
        }
    })

}
