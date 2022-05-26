//  Copyright (c) 2022, DURUYAO. All rights reserved.
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
