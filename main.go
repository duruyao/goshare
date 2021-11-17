package main

import (
	"flag"
	"fmt"
	"os"
)

var quit = make(chan struct{})
var servAddr = flag.String("a", "127.0.0.1:8080", "listening address in \"<ip>:<port>\" format")
var filePath = flag.String("f", UserHomeDir(), "handling local file path in \"/.../<path>\" format")

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			fmt.Println("GoFS version 2021.11.16")
			return
		}
	}
	flag.Parse()
	fmt.Printf("GoFS is listening on %s and handling %s ...\n", *servAddr, *filePath)
	go GoRunWebApp(*servAddr, *filePath)
	<-quit
}
