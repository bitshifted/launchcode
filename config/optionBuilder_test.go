package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetModulePathMac(t *testing.T) {
	ismacOsFunc = func() bool {
		return true
	}
	modulePath = "modules"
	result := setModulePath()
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "--module-path", result[0])
	assert.Equal(t, "Contents/Runtime/modules", result[1])
}

func Test_SetModulePathNonMac(t *testing.T) {
	ismacOsFunc = func() bool {
		return false
	}
	modulePath = "modules"
	result := setModulePath()
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "--module-path", result[0])
	assert.Equal(t, "modules", result[1])
}

func Test_SetClasspathPathMac(t *testing.T) {
	ismacOsFunc = func() bool {
		return true
	}
	classpath = "cp/*"
	result := setClasspath()
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "-classpath", result[0])
	assert.Equal(t, "Contents/Runtime/cp/*", result[1])
}

func Test_SetClasspathPathNonMac(t *testing.T) {
	ismacOsFunc = func() bool {
		return false
	}
	classpath = "cp/*"
	result := setClasspath()
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "-classpath", result[0])
	assert.Equal(t, "cp/*", result[1])
}

func Test_SetSplashScreenMac(t *testing.T) {
	ismacOsFunc = func() bool {
		return true
	}
	splashScreen = "-splash:test-splash.png"
	result := setSplashScreen()
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "-splash:Contents/Resources/test-splash.png", result[0])
}
