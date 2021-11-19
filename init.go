package main

import "flag"

var (
	ShowVersion bool
	ListenAddr string
	HandlePath string
	UrlPrefix  string
)

const (
	Version          = "GoFS Version 2021.11.18"
	UrlPrefixDefault = "Base of file path"
)

func init() {
	flag.BoolVar(&ShowVersion, "version", false, "version of GoFS")
	flag.StringVar(&UrlPrefix, "prefix", UrlPrefixDefault, "prefix of url path")
	flag.StringVar(&ListenAddr, "addr", "127.0.0.1:8080", "listening address in \"ip:port\" format")
	flag.StringVar(&HandlePath, "path", UserHomeDir(), "handling file path in \"/path/to/file\" format")
	flag.Parse()
}
