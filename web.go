package main

import (
	"log"
	"net/http"
	"time"
)

// GoRunWebApp starts files transfer service.
func GoRunWebApp(listenAddr string, handleDir string, urlPrefix string) {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(handleDir))
	mux.Handle("/"+urlPrefix+"/", http.StripPrefix("/"+urlPrefix+"/", fileHandler))

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
