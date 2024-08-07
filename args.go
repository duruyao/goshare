//  Copyright 2022-2032 Ryan Du <duruyao@gmail.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

const (
	DefaultHost      = `0.0.0.0:80`
	DefaultScheme    = `http`
	DefaultUrlPrefix = `/`
	UsageTmpl        = `{{.Logo}}
Usage: {{.Exec}} [OPTIONS]

GoShare shares files and directories via HTTP protocol

Options:
    -h, --help                  Display this help message
    --host STRING               Host address to listen (default: '{{.Host}}')
    --path STRING               Path or directory (default: '{{.Path}}')
    --scheme STRING             Scheme name (default: '{{.Scheme}}')
    --url-prefix STRING         Custom URL prefix (default: '{{.UrlPrefix}}')
    -v, --version               Print version information and quit

Examples:
    {{.Exec}} -host {{.Host}} -path {{.ExamplePath}}
    {{.Exec}} --host {{.Host}} --path {{.ExamplePath}} --url-prefix /{{.ExampleUrlPrefix}}
    {{.Exec}} --host={{.Host}} --path={{.ExamplePath}} --url-prefix=/{{.ExampleUrlPrefix}}

See more about {{.App}} at {{.Link}}
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
		Logo             string
		Exec             string
		Host             string
		Path             string
		Scheme           string
		UrlPrefix        string
		ExamplePath      string
		ExampleUrlPrefix string
		App              string
		Link             string
	}{
		Logo:             Logo,
		Exec:             os.Args[0],
		Host:             DefaultHost,
		Path:             DefaultPath,
		Scheme:           DefaultScheme,
		UrlPrefix:        DefaultUrlPrefix,
		ExamplePath:      CurrentDirMust(),
		ExampleUrlPrefix: filepath.Base(CurrentDirMust()),
		App:              App,
		Link:             Link,
	}
	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
