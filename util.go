package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func ParseFromPath(handlePath string) (handleDir string, urlPath string) {
	if info, err := os.Stat(handlePath); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", handlePath)
	} else {
		handlePath, _ = filepath.Abs(handlePath)
		if info.IsDir() {
			handleDir = handlePath
			urlPath = ""
			//urlPath = filepath.Base(handleDir)
		} else {
			handleDir = filepath.Dir(handlePath)
			//urlPath = filepath.Base(handleDir) + "/" + filepath.Base(handlePath)
			urlPath = filepath.Base(filepath.Base(handlePath))
		}
	}
	return handleDir, urlPath
}
