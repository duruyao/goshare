package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	arg  = NewArgument()
	quit = make(chan struct{})
)

func main() {
	if arg.WantHelp() {
		fmt.Println(arg.Usage())
		return
	}
	if arg.WantVersion() {
		fmt.Println(VersionSerial())
		return
	}

	host := arg.Host()
	path := AbsPathMust(arg.Path())
	dir, file := "", ""
	scheme := arg.Scheme()
	urlPrefix := FixedUrlPrefix(arg.UrlPrefix())

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Error: No such file or directory: " + path)
		fmt.Println(arg.Usage())
		return
	}
	if info.IsDir() {
		dir, file = path, ""
	} else {
		dir, file = filepath.Dir(path), filepath.Base(path)
	}
	fmt.Println(dir)
	go StartHttpFileService(host, dir, urlPrefix)
	fmt.Printf("Share files by the URL %s://%s%s%s\n", scheme, host, urlPrefix, file)
	<-quit
}
