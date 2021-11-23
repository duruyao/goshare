package main

import (
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
