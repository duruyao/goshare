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
	AppLink           = `https://github.com/duruyao/goshare`
	Version           = `1.0.0`
	ReleaseDate       = `2022-05-23`
	VersionSerialTmpl = `{{.App}} {{.Version}} ({{.ReleaseDate}})`
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
