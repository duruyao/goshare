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
		fmt.Println(err.Error())
		fmt.Println(arg.Usage())
		return
	}
	if info.IsDir() {
		dir, file = path, ""
	} else {
		dir, file = filepath.Dir(path), filepath.Base(path)
	}
	go StartHttpFileService(host, dir, urlPrefix)
	fmt.Println(RunningStatus(dir, host, scheme, urlPrefix, file))
	<-quit
}
