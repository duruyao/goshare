package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	WantVersion   bool
	Scheme        string
	UrlPrefix     string
	HandlingPath  string
	ListeningAddr string
)

const (
	SH            = ` (shorthand)`
	VersionSerial = `GoShare Version 2021.12.22`
	UsageTmpl     = `Usage: {{.App}} [OPTIONS] PATH

The tool share files or directories by HTTP or FTP protocol

Options:
	-d, --detach				Run service in background
    -h, --help      			Display this help message
    --host STRING          		Host address to listen (default: '{{.DefaultAddr}}')
    --scheme STRING				Scheme name (default: '{{.DefaultScheme}}')
    --url-prefix STRING 		URL prefix (default: 'PATH')
    -v, --version   			Print version information and quit

Examples:
    {{.App}} -host example.com /opt/share0/releases/
    {{.App}} -host localhost:3927 /opt/share0/releases/
    {{.App}} --host localhost:3927 --url-prefix /releases/ /opt/share0/releases/
    {{.App}} --host=localhost:3927 --url-prefix=/releases/ /opt/share0/releases/
`
)

func WantUsage() bool {
	helpArgs := map[string]bool{
		"-h":     true,
		"--h":    true,
		"-help":  true,
		"--help": true,
	}
	for _, arg := range os.Args[1:] {
		if helpArgs[arg] {
			return true
		}
	}
	return false
}

func ShowUsage() {
	tmpl := template.Must(template.New("usage tmpl").Parse(UsageTmpl))

	data := struct {
		App           string
		DefaultScheme string
		DefaultAddr   string
		DefaultPath   string
	}{
		App:           os.Args[0],
		DefaultPath:   UserHomeDirMust(),
		DefaultScheme: "http",
		DefaultAddr:   "localhost:3927",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}

func ShowVersion() {
	fmt.Println(VersionSerial)
}

func InitArgs() {
	flag.StringVar(&UrlPrefix, "url-prefix", "", "url prefix")

	flag.StringVar(&Scheme, "s", "http", "scheme name"+SH)
	flag.StringVar(&Scheme, "scheme", "http", "scheme name")

	flag.BoolVar(&WantVersion, "v", false, "show version"+SH)
	flag.BoolVar(&WantVersion, "version", false, "show version")

	flag.StringVar(&HandlingPath, "p", UserHomeDirMust(), "handing path or directory"+SH)
	flag.StringVar(&HandlingPath, "path", UserHomeDirMust(), "handing path or directory")

	flag.StringVar(&ListeningAddr, "a", "127.0.0.1:8080", "listening address"+SH)
	flag.StringVar(&ListeningAddr, "address", "127.0.0.1:8080", "listening address")
}

// ParseArgs parses some values from arguments.
func ParseArgs() (addr string, dir string, filename string, prefix string, url string, err error) {
	if WantUsage() {
		ShowUsage()
		err = errors.New("just show usage")
		return
	}

	flag.Parse()
	if WantVersion {
		ShowVersion()
		err = errors.New("just show version")
		return
	}

	info, e := os.Stat(HandlingPath)
	if os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", HandlingPath)
		err = e
		return
	}

	addr = ListeningAddr // NOTE: set addr
	HandlingPath = AbsPathMust(HandlingPath)
	if info.IsDir() {
		dir = HandlingPath
		filename = ""
	} else {
		dir = filepath.Dir(HandlingPath)       // NOTE: set dir
		filename = filepath.Base(HandlingPath) // NOTE: set filename
	}

	if UrlPrefix != "" { // FIXME: support url prefix = /
		if UrlPrefix[0] == '/' {
			UrlPrefix = UrlPrefix[1:]
		}
		if UrlPrefix[len(UrlPrefix)-1] == '/' {
			UrlPrefix = UrlPrefix[:len(UrlPrefix)-1]
		}
		prefix = UrlPrefix
	} else {
		prefix = filepath.Base(dir) // NOTE: set prefix
	}

	url = fmt.Sprintf("%s://%s/%s/%s", "http", addr, prefix, filename) // NOTE: set url

	return
}
