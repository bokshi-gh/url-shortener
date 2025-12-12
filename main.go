package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"url-shortner/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	// http.HandleFunc("/login", handlers.LoginHandler)
	// http.HandleFunc("/delete-account", handlers.DeleteAccountHandler)

	// http.HandleFunc("/add-url", handlers.AddURLHandler)
	// http.HandleFunc("/delete-url", handlers.DeleteURLHandler)
	// http.HandleFunc("/r/", handlers.RedirectURLHandler)

	fs := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("./static", r.URL.Path)

		f, err := os.Open(path)
		if err != nil {
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		defer f.Close()

		stat, _ := f.Stat()
		if stat.IsDir() {
			http.ServeFile(w, r, "./static/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	})
	

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

