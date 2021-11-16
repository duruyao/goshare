package main

import (
	"log"
	"net/http"
	"time"
)

// GoRunWebApp start up a file server.
func GoRunWebApp(servAddr string, filePath string) {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(filePath))
	mux.Handle("/", http.StripPrefix("/", fileHandler))

	app := http.Server{
		Addr:              servAddr,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       7 * time.Second,
		ReadHeaderTimeout: 7 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       7 * time.Second,
		MaxHeaderBytes:    1024 * 1024,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          log.Default(),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	log.Fatalln(app.ListenAndServe())
}
