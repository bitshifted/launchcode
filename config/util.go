package config

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	macOsDirPath = "Contents/MacOS"
	jreDirPath   = "Contents/Resources"
)

func GetAppBaseDirectory(exePath string) string {
	appDir := filepath.Dir(exePath)
	// special handling in Mac OS X
	if runtime.GOOS == "darwin" {
		log.Println("Adjusting app base directory for Mac OS X")
		appDir = strings.Replace(exePath, macOsDirPath, "", 1)
	}
	return appDir
}
