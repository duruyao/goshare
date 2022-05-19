package main

import (
	"log"
	"net/http"
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

	log.Fatalln(app.ListenAndServe())
}
