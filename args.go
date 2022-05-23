package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
	"sync"
	"text/template"
)

const (
	DefaultHost      = `localhost:3927`
	DefaultScheme    = `http`
	DefaultUrlPrefix = `/`
	UsageTmpl        = `Usage: {{.Exec}} [OPTIONS]

GoShare shares file and directory by HTTP or FTP protocol

Options:
    -h, --help                  Display this help message
    --host STRING               Host address to listen (default: '{{.Host}}')
    --path STRING               Path or directory (default: '{{.Path}}')
    --scheme STRING             Scheme name (default: '{{.Scheme}}')
    --url-prefix STRING         Custom URL prefix (default: '{{.UrlPrefix}}')
    -v, --version               Print version information and quit

Examples:
    {{.Exec}} -host example.io -path /opt/share0/releases/
    {{.Exec}} -host {{.Host}} -path /opt/share0/releases/
    {{.Exec}} --host {{.Host}} --url-prefix /releases/ --path /opt/share0/releases/
    {{.Exec}} --host={{.Host}} --url-prefix=/releases/ --path=/opt/share0/releases/

See more about {{.App}} at {{.AppLink}}
`
)

var (
	DefaultPath = CurrentDirMust()
)

type argument struct {
	WantHelp    bool   `json:"want_help"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	Scheme      string `json:"scheme"`
	UrlPrefix   string `json:"url_prefix"`
	WantVersion bool   `json:"want_version"`
}

type Argument struct {
	arg       *argument
	parseOnce sync.Once
}

func NewArgument() *Argument {
	return &Argument{arg: &argument{}}
}

func (a *Argument) Init() {
	flag.BoolVar(&(a.arg.WantHelp), "h", false, "Display this help message")
	flag.BoolVar(&(a.arg.WantHelp), "help", false, "Display this help message")
	flag.StringVar(&(a.arg.Host), "host", DefaultHost, "Host address to listen")
	flag.StringVar(&(a.arg.Path), "path", DefaultPath, "Path or directory")
	flag.StringVar(&(a.arg.Scheme), "scheme", DefaultScheme, "Scheme name")
	flag.StringVar(&(a.arg.UrlPrefix), "url-prefix", DefaultUrlPrefix, "Custom URL prefix")
	flag.BoolVar(&(a.arg.WantVersion), "v", false, "Print version information and quit")
	flag.BoolVar(&(a.arg.WantVersion), "version", false, "Print version information and quit")
}

func (a *Argument) Parse() {
	a.parseOnce.Do(func() {
		flag.Parse()
	})
}

func (a *Argument) Serialize() ([]byte, error) {
	return json.Marshal(a.arg)
}

func (a *Argument) Deserialize(js []byte) error {
	return json.Unmarshal(js, a.arg)
}

func (a *Argument) String() string {
	a.Parse()
	if js, err := json.MarshalIndent(a.arg, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (a *Argument) WantHelp() bool {
	a.Parse()
	return a.arg.WantHelp
}

func (a *Argument) Host() string {
	a.Parse()
	return a.arg.Host
}

func (a *Argument) Path() string {
	a.Parse()
	return a.arg.Path
}

func (a *Argument) Scheme() string {
	a.Parse()
	return a.arg.Scheme
}

func (a *Argument) UrlPrefix() string {
	a.Parse()
	return a.arg.UrlPrefix
}

func (a *Argument) WantVersion() bool {
	a.Parse()
	return a.arg.WantVersion
}

func (a *Argument) Usage() string {
	a.Parse()
	tmpl := template.Must(template.New("usage tmpl").Parse(UsageTmpl))
	data := struct {
		App       string
		AppLink   string
		Exec      string
		Host      string
		Path      string
		Scheme    string
		UrlPrefix string
	}{
		App:       App,
		AppLink:   AppLink,
		Exec:      os.Args[0],
		Host:      DefaultHost,
		Path:      DefaultPath,
		Scheme:    DefaultScheme,
		UrlPrefix: DefaultUrlPrefix,
	}
	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
