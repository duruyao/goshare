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
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

// AbsPathMust returns the absolute path of the given path.
func AbsPathMust(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}
	return abs
}

var (
	userHomeDirOnce sync.Once
	userHomeDir     string
)

// UserHomeDirMust returns the home directory of the current user, such as "/home/foo" in Unix-like OS.
func UserHomeDirMust() string {
	userHomeDirOnce.Do(func() {
		var err error
		userHomeDir, err = os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
	})
	return userHomeDir
}

var (
	currentDirOnce sync.Once
	currentDir     string
)

// CurrentDirMust returns the current working directory.
func CurrentDirMust() string {
	currentDirOnce.Do(func() {
		var err error
		currentDir, err = os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
	})
	return currentDir
}

const (
	App               = `GoShare`
	Link              = `https://github.com/duruyao/goshare`
	Version           = `1.2.0`
	ReleaseDate       = `2024-07-29`
	VersionSerialTmpl = `{{.App}} {{.Version}} ({{.ReleaseDate}})`
	Logo              = `
   _____       _____ _
  / ____|     / ____| |
 | |  __  ___| (___ | |__   __ _ _ __ ___
 | | |_ |/ _ \\___ \| '_ \ / _' | '__/ _ \
 | |__| | (_) |___) | | | | (_| | | |  __/
  \_____|\___/_____/|_| |_|\__,_|_|  \___|

`
)

// VersionSerial returns version serial.
func VersionSerial() string {
	tmpl := template.Must(template.New("version serial tmpl").Parse(VersionSerialTmpl))
	data := struct {
		App         string
		Version     string
		ReleaseDate string
	}{
		App:         App,
		Version:     Version,
		ReleaseDate: ReleaseDate,
	}
	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}

// FixedUrlPrefix returns the result after fixing the given URL prefix.
func FixedUrlPrefix(urlPrefix string) string {
	if "/" == urlPrefix {
		return "/"
	}
	return AbsPathMust("/"+urlPrefix) + "/"
}

const (
	RunningStatusTmpl = `{{.Logo}}
{{.App}} is handling directory '{{.Dir}}' and listening on '{{.Host}}'

Access your shared files via this URL {{.Scheme}}://{{.Host}}{{.UrlPrefix}}{{.File}}
`
)

// RunningStatus returns server running status.
func RunningStatus(dir string, host string, scheme string, urlPrefix string, file string) string {
	tmpl := template.Must(template.New("running status tmpl").Parse(RunningStatusTmpl))
	data := struct {
		Logo      string
		App       string
		Dir       string
		Host      string
		Scheme    string
		UrlPrefix string
		File      string
	}{
		Logo:      Logo,
		App:       App,
		Dir:       dir,
		Host:      host,
		Scheme:    scheme,
		UrlPrefix: urlPrefix,
		File:      file,
	}
	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
