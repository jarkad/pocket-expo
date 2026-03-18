// This repo is a simple example of using Templ, Gorm, and net/http.
package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/handlers"
	"git.jarkad.net.eu.org/jarkad/pocket-expo/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed static
var assets embed.FS

//go:generate templ generate

func main() {
	log := log.Default()

	db, err := gorm.Open(sqlite.Open(
		"db.sqlite3?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"))
	if err != nil {
		log.Fatalln("cannot open database:", err)
	}

	counter, err := services.NewCounter(db)
	if err != nil {
		log.Fatalln("while initializing counter:", err)
	}

	homepageHandler := handlers.NewHomepageHandler(log, counter)

	mux := http.NewServeMux()
	mux.Handle("GET  /{$}", homepageHandler)
	mux.Handle("POST /{$}", homepageHandler)
	mux.Handle("GET  /static/", http.FileServerFS(assets))

	handler := http.Handler(mux)
	handler = withLogging(log)(handler)

	server := http.Server{ //nolint:exhaustruct
		Addr:              ":3000",
		Handler:           handler,
		ReadTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 1 * time.Minute,
		WriteTimeout:      1 * time.Minute,
		IdleTimeout:       1 * time.Minute,
		ErrorLog:          log,
	}

	log.Println("Listening on", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("could not listen:", err)
	}
}

func withLogging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println(r.Method, r.URL.String())
			h.ServeHTTP(w, r)
		})
	}
}
