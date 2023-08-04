package main

import (
	"github.com/rednafi/shorturl/src"
	"net/http"
)

func main() {
	http.HandleFunc("/", src.DisplayIndex) 	// Display the home page
	http.HandleFunc("/shorten/", src.ShortenUrl) // Shorten the URL
	http.HandleFunc("/r/", src.RedirectUrl) 	// Redirect to the original URL
	http.ListenAndServe(":8080", nil)

	defer src.Db.Close()
}
