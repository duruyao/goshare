package main

import (
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// GoRunWebApp start up a file server.
func GoRunWebApp(listenAddr string, handleDir string) {
	mux := http.NewServeMux()
	//fileHandler1 := http.FileServer(http.Dir(handleDir))
	//mux.Handle("/", http.StripPrefix("/", fileHandler1))

	fileHandler2 := http.FileServer(http.Dir(handleDir))
	prefix := "/" + filepath.Base(handleDir) + "/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileHandler2))

	app := http.Server{
		Addr:              listenAddr,
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
