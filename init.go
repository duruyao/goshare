package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

var (
	ShowVersion   bool
	Scheme        string
	UrlPrefix     string
	HandlingPath  string
	ListeningAddr string
)

const (
	Version   = "GoFS Version 2021.11.18"
	UsageTmpl = `USAGE:
    {{.AppPath}} [-v] [-h]
            [--url-prefix <prefix>]
            [-s {http, https, ftp}] [-a <address>] [-p <path>]

OPTIONS:
    -v, --version
                    show version
    -h, --help
                    show help manual
    --url-prefix <prefix>
                    url prefix
    -s {http, https, ftp}, --scheme {http, https, ftp}
                    scheme name (default: "{{.DefaultScheme}}")
    -a <ip:port>, --address <ip:port>
                    listening address (default: "{{.DefaultAddr}}")
    -p </path/to/file>,	--path </path/to/file>
                    handing path or directory (default: "{{.DefaultPath}}")`
)

type UsageStruct struct {
	AppPath       string
	DefaultScheme string
	DefaultAddr   string
	DefaultPath   string
}

func init() {
	flag.StringVar(&UrlPrefix, "prefix", "", "url prefix")
	flag.StringVar(&Scheme, "scheme", "http", "scheme name")
	flag.StringVar(&HandlingPath, "path", UserHomeDirMust(), "handling path or directory")
	flag.StringVar(&ListeningAddr, "addr", "127.0.0.1:8080", "listening address")
	flag.BoolVar(&ShowVersion, "version", false, "show version of GoFS")
	flag.Parse()
}

func ShowUsage() {
	tmpl := template.Must(template.New("usage tmpl").Parse(UsageTmpl))

	data := UsageStruct{
		AppPath:       os.Args[0],
		DefaultPath:   UserHomeDirMust(),
		DefaultScheme: "http",
		DefaultAddr:   "127.0.0.1:8080",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}
