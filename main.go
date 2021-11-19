package main

import (
	"fmt"
)

var quit = make(chan struct{})

func main() {
	addr, dir, name, prefix, url, err := ParseArgs()
	if err != nil {
		return
	}
	fmt.Printf("GoFS is listening on %s and handling %s/%s ...\n", addr, dir, name)
	go GoRunWebApp(addr, dir, prefix)
	fmt.Printf("Access or share the URL: %s\n",  url)
	<-quit
}
