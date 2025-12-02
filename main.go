package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"url-shortner/api"
)

func main() {
	http.HandleFunc("/api/register", api.RegisterHandler)
	http.HandleFunc("/api/login", api.LoginHandler)
	// http.HandleFunc("/api/change-username", api.ChangeUsernameHandler)
	// http.HandleFunc("/api/change-password", api.ChangePasswordHandler)
	http.HandleFunc("/api/delete-account", api.DeleteAccountHandler)

	http.HandleFunc("/api/add-url", api.AddURLHandler)
	// http.HandleFunc("/api/change-url", api.ChangeURLHandler)
	http.HandleFunc("/api/delete-url", api.DeleteURLHandler)


	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("./static", r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// If file doesn't exist, serve index.html (SPA support)
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	}))

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

