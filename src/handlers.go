package src

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		panic(err)
	}

	// Create table and indexes if they don't exist
	stmt, err := Db.Prepare(
		`CREATE TABLE IF NOT EXISTS url_mappings (
		id TEXT,
		hash TEXT,
		url TEXT,
		PRIMARY KEY (id),
		UNIQUE (hash)
	);

	CREATE INDEX IF NOT EXISTS idx_url_mappings_hash ON url_mappings (hash);`)
	if err != nil {
		panic(err)
	}
	stmt.Exec()
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)

	// Check if method is POST and return error if not
	if r.Method != http.MethodPost {
		JsonError(
			w,
			map[string]string{"error": "Method not allowed"},
			http.StatusMethodNotAllowed,
		)
	}

	// Check if decoding was successful
	if err != nil {
		log.Println(err)
		JsonError(
			w,
			map[string]string{"error": "Decoding error"},
			http.StatusBadRequest,
		)
		return
	}

	// Validate url
	if !ValidateUrl(req.Url) {
		log.Println("Invalid URL")
		JsonError(
			w,
			map[string]string{"error": "Invalid URL"},
			http.StatusBadRequest,
		)
		return
	}

	// Strip trailing slash
	req.Url = TrimSlash(req.Url)

	// Generate hash
	hash := GenerateHash(req.Url)

	// Check if hash exists and return id if it does
	row := Db.QueryRow("SELECT id FROM url_mappings WHERE hash=?", hash)
	var id string
	err = row.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		JsonError(
			w,
			map[string]string{"error": "Error querying database"},
			http.StatusInternalServerError,
		)
		return
	}

	// If id exists, return it as tinyurl and exit
	if id != "" {
		res := Response{TinyUrl: id}
		jsonRes, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
		return
	}

	// If id doesn't exist, generate a new one
	id = GenerateId(7)

	// Insert new mapping
	_, err = Db.Exec("INSERT INTO url_mappings VALUES (?, ?, ?)", id, hash, req.Url)
	if err != nil {
		log.Println(err)
		JsonError(
			w,
			map[string]string{"error": "Error inserting into database"},
			http.StatusInternalServerError,
		)
		return
	}

	// Return new ID
	res := Response{TinyUrl: id}
	jsonRes, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	// Get the short URL from the request path
	id := TrimSlash(r.URL.Path)

	// Query the database for the long URL
	var url string
	err := Db.QueryRow("SELECT url FROM url_mappings WHERE id = ?", id).Scan(&url)

	// If the short URL is not found in the database, return a 404 error
	if err != nil {
		log.Println(err)
		JsonError(w, map[string]string{"error": "Not found"}, http.StatusNotFound)
		return
	}

	// Redirect to the long URL
	http.Redirect(w, r, url, http.StatusSeeOther)
}
