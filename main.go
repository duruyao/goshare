package main

import (
	"flag"
	"fmt"
)

var quit = make(chan struct{})
var servAddr = flag.String("a", "127.0.0.1:8080", "listening address in \"<ip>:<port>\" format")
var filePath = flag.String("f", UserHomeDir(), "handling local file path in \"/.../<path>\" format")

func main() {
	flag.Parse()
	fmt.Printf("GoFS is listening on %s and handling %s ...\n", *servAddr, *filePath)
	go GoRunWebApp(*servAddr, *filePath)
	<-quit
}
