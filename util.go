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

// UserHomeDir returns the home directory of the current user, such as "/home/user" in Unix-like OS.
func UserHomeDir() string {
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
func ParseArgs() (addr string, dir string, name string, prefix string, url string, err error) {
	if ShowVersion {
		fmt.Println(Version)
		err = errors.New("show Version")
	} else if info, e := os.Stat(HandlePath); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", HandlePath)
		err = e // NOTE: set err
	} else {
		addr = ListenAddr // NOTE: set addr

		HandlePath, _ = filepath.Abs(HandlePath)
		if info.IsDir() {
			dir = HandlePath
			name = ""
		} else {
			dir = filepath.Dir(HandlePath) // NOTE: set dir
			name = filepath.Base(HandlePath) // NOTE: set name
		}

		if UrlPrefix != UrlPrefixDefault {
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

		url = fmt.Sprintf("%s://%s/%s/%s", "http", addr, prefix, name) // NOTE: set url
	}
	return
}
