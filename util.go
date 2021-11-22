package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	userHomeDirOnce sync.Once
	userHomeDir     string
)

func AbsPathMust(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}
	return abs
}

// UserHomeDirMust returns the home directory of the current user, such as "/home/user" in Unix-like OS.
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

// ParseArgs parses some values from arguments.
func ParseArgs() (addr string, dir string, filename string, prefix string, url string, err error) {
	if ShowVersion {
		fmt.Println(Version)
		err = errors.New("just show version")
	} else if info, e := os.Stat(HandlingPath); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", HandlingPath)
		err = e // NOTE: set err
	} else {
		addr = ListeningAddr // NOTE: set addr

		HandlingPath, _ = filepath.Abs(HandlingPath)
		if info.IsDir() {
			dir = HandlingPath
			filename = ""
		} else {
			dir = filepath.Dir(HandlingPath)       // NOTE: set dir
			filename = filepath.Base(HandlingPath) // NOTE: set filename
		}

		if UrlPrefix != "" {
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
	}
	return
}
