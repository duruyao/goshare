package main

import (
	"flag"
	"fmt"
	"os"
)

var quit = make(chan struct{})
var listenAddr = flag.String("a", "127.0.0.1:8080", "listening address in \"<ip>:<port>\" format")
var handlePath = flag.String("f", UserHomeDir(), "handling local file path in \"/.../<path>\" format")

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			fmt.Println("GoFS version 2021.11.16")
			return
		}
	}
	flag.Parse()
	handleDir, urlPath := ParseFromPath(*handlePath)
	if handleDir == "" {
		return
	}
	fmt.Printf("GoFS is listening on %s and handling %s ...\n", *listenAddr, *handlePath)
	go GoRunWebApp(*listenAddr, handleDir)
	fmt.Printf("Access %s://%s/%s\n", "http", *listenAddr, urlPath)
	<-quit
}
