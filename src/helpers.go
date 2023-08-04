package src

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

// ValidateUrl checks if a given URL is valid
func ValidateUrl(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false
	}
	u, err := url.Parse(rawUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

// Strip trailing slash from url
func TrimSlash(rawUrl string) string {
	// Trim prefix
	rawUrl = strings.TrimPrefix(rawUrl, "/")
	rawUrl = strings.TrimSuffix(rawUrl, "/")
	return rawUrl
}

// Generate a random string of length n
func GenerateId(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}

// Generate hash of the url
func GenerateHash(rawUrl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(rawUrl))
	return hex.EncodeToString(hasher.Sum(nil))
}

// JSONError writes an error to the response
func JsonError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

// GetQualifiedTinyUrl returns the qualified tinyurl
func GetQualifiedTinyUrl(r *http.Request, id string) string {
	if r.TLS != nil {
		return fmt.Sprintf("https://%s/r/%s", r.Host, id)
	}
	return fmt.Sprintf("http://%s/r/%s", r.Host, id)
}
