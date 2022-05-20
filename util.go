package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

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

// CurrentDirMust returns the current directory.
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
