package main

import (
	"log"
	"os"
	"sync"
)

var userHomeDirOnce sync.Once
var userHomeDir string

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
