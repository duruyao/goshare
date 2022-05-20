package main

import (
	"log"
	"net/http"
)

// StartHttpFileService starts a file service using HTTP protocol.
func StartHttpFileService(host string, dir string, prefix string) {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(dir))
	mux.Handle(prefix, http.StripPrefix(prefix, fileHandler))

	server := http.Server{
		Addr:              host,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          log.Default(),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	log.Fatalln(server.ListenAndServe())
}
