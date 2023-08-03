package main

import (
	"github.com/rednafi/shorturl/src"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", src.ShortenUrl)
	http.HandleFunc("/", src.RedirectUrl)
	http.ListenAndServe(":8080", nil)
	defer src.Db.Close()
}
