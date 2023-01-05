package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	macOsDirPath       = "/Contents/MacOS"
	jreDirPathMac      = "Contents/Runtime/jre"
	runtimeDirPathMac  = "Contents/Runtime"
	macOsResourcesPath = "Contents/Resources"
	goOsMac            = "darwin"
)

func isMacOs() bool {
	return runtime.GOOS == goOsMac
}

func GetAppBaseDirectory(exePath string) string {
	appDir := filepath.Dir(exePath)
	// special handling in Mac OS X
	if ismacOsFunc() {
		log.Println("Adjusting app base directory for Mac OS X")
		appDir = strings.Replace(appDir, macOsDirPath, "", 1)
	}
	return appDir
}

func createPathWithPrefix(prefixPath, path string) string {
	prefix := ""
	if ismacOsFunc() {
		prefix = prefixPath + "/"
	}
	return fmt.Sprintf("%s%s", prefix, strings.TrimSpace(path))
}
