//  Copyright 2022-2032 Ryan Du <duruyao@gmail.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// HttpStaticFS starts a file service using HTTP protocol.
func HttpStaticFS(host string, dir string, prefix string) {
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

// HttpStaticFile handles a static file using HTTP protocol.
func HttpStaticFile(host string, filename string, prefix string) {
	mux := http.NewServeMux()
	fileHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filename)
		if err != nil {
			http.Error(w, "No such file: "+filename, http.StatusNotFound)
			return
		}
		defer func() { _ = file.Close() }()
		info, err := file.Stat()
		if err != nil {
			http.Error(w, "Error getting file info: "+filename, http.StatusInternalServerError)
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file: "+filename, http.StatusInternalServerError)
			return
		}
		contentType := mime.TypeByExtension(filepath.Ext(filename))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
		_, _ = w.Write(data)
	})
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
