package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppBaseDirMacOS(t *testing.T) {
	ismacOsFunc = func() bool {
		return true
	}
	exePath := "/Applications/MyApp.app/Contents/MacOS/binary"
	result := GetAppBaseDirectory(exePath)
	assert.Equal(t, "/Applications/MyApp.app", result)
}

func TestCreatePathWithPrefixtMacOs(t *testing.T) {
	ismacOsFunc = func() bool {
		return true
	}
	result := createPathWithPrefix(jreDirPathMac, "modules")
	assert.Equal(t, jreDirPathMac+"/modules", result)
	result = createPathWithPrefix(macOsResourcesPath, "file.txt")
	assert.Equal(t, macOsResourcesPath+"/file.txt", result)
}

func TestCreatePathWithPrefixtNonMacOs(t *testing.T) {
	ismacOsFunc = func() bool {
		return false
	}
	result := createPathWithPrefix(jreDirPathMac, "modules")
	assert.Equal(t, "modules", result)
	result = createPathWithPrefix(macOsResourcesPath, "file.txt")
	assert.Equal(t, "file.txt", result)
}
