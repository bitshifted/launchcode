package main

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"gopkg.in/yaml.v3"
)

const (
	defaultOutputDir = "output"
	linuxOutputDir   = "linux"
	macOutputDir     = "mac"
	windowsOutputDir = "windows"
	embedJvmDir      = "jvm-dir"
	embedJvmOptions  = "jvmopts.txt"
	embedJvmProps    = "jvmprops.txt"
	embedSplash      = "splash.txt"
	embedModulePath  = "modulepath.txt"
	embedAddModules  = "addmodules.txt"
	embedModule      = "module.txt"
	embedClasspath   = "classpath.txt"
	embedMainClass   = "mainclass.txt"
	embedJar         = "jar.txt"
	embedArgs        = "args.txt"
	shellRegex       = `\s+\\\s+`
)

type JavaOptions struct {
	JreDirectory        string `yaml:"jre-dir"`
	RestartCode         int    `yaml:"restart-code"`
	JvmOptions          string `yaml:"jvm-options"`
	JvmSystemProperties string `yaml:"jvm-system-properties"`
	SplashScreen        string `yaml:"splash-screen"`
	ModulePath          string `yaml:"module-path"`
	AddModules          string `yaml:"add-modules"`
	Module              string `yaml:"module"`
	Classpath           string `yaml:"classpath"`
	MainClass           string `yaml:"main-class"`
	Jar                 string `yaml:"jar"`
	Arguments           string `yaml:"arguments"`
}

func (jo *JavaOptions) copyFrom(source JavaOptions) {
	if source.JreDirectory != "" {
		jo.JreDirectory = source.JreDirectory
	}
	if source.RestartCode != 0 {
		jo.RestartCode = source.RestartCode
	}
	if source.JvmOptions != "" {
		jo.JvmOptions = source.JvmOptions
	}
	if source.JvmSystemProperties != "" {
		jo.JvmSystemProperties = source.JvmSystemProperties
	}
	if source.SplashScreen != "" {
		jo.SplashScreen = source.SplashScreen
	}
	if source.ModulePath != "" {
		jo.ModulePath = source.ModulePath
	}
	if source.AddModules != "" {
		jo.AddModules = source.AddModules
	}
	if source.Module != "" {
		jo.Module = source.Module
	}
	if source.Classpath != "" {
		jo.Classpath = source.Classpath
	}
	if source.MainClass != "" {
		jo.MainClass = source.MainClass
	}
	if source.Jar != "" {
		jo.Jar = source.Jar
	}
	if source.Arguments != "" {
		jo.Arguments = source.Arguments
	}
}

func (jo *JavaOptions) writeConfig(directory string) error {
	reg := regexp.MustCompile(shellRegex)
	err := os.WriteFile(path.Join(directory, embedJvmDir), []byte(jo.JreDirectory), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedJvmOptions), []byte(reg.ReplaceAllString(jo.JvmOptions, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedJvmProps), []byte(reg.ReplaceAllString(jo.JvmSystemProperties, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedJvmProps), []byte(reg.ReplaceAllString(jo.JvmSystemProperties, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedSplash), []byte(reg.ReplaceAllString(jo.SplashScreen, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedModulePath), []byte(reg.ReplaceAllString(jo.ModulePath, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedAddModules), []byte(reg.ReplaceAllString(jo.AddModules, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedModule), []byte(reg.ReplaceAllString(jo.Module, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedClasspath), []byte(reg.ReplaceAllString(jo.Classpath, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedMainClass), []byte(reg.ReplaceAllString(jo.MainClass, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedJar), []byte(reg.ReplaceAllString(jo.Jar, " ")), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(directory, embedArgs), []byte(reg.ReplaceAllString(jo.Arguments, " ")), 0644)
	if err != nil {
		return err
	}

	return err
}

type LauncherConfig struct {
	Common  JavaOptions `yaml:"common"`
	Linux   JavaOptions `yaml:"linux"`
	Windows JavaOptions `yaml:"windows"`
	Mac     JavaOptions `yaml:"mac"`
}

func (lc *LauncherConfig) printConfig() error {
	// linux config
	linuxOutputPath := path.Join(defaultOutputDir, linuxOutputDir)
	err := os.MkdirAll(linuxOutputPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = lc.Linux.writeConfig(linuxOutputPath)
	if err != nil {
		return err
	}
	// windows config
	windowsOutputPath := path.Join(defaultOutputDir, windowsOutputDir)
	err = os.MkdirAll(windowsOutputPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = lc.Windows.writeConfig(windowsOutputPath)
	if err != nil {
		return err
	}
	// mac config
	macutputPath := path.Join(defaultOutputDir, macOutputDir)
	err = os.MkdirAll(macutputPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = lc.Mac.writeConfig(macutputPath)
	return err
}

func main() {
	if len(os.Args) < 2 {
		panic("Configuration file not specified")
	}
	configFilePath := os.Args[1]
	if configFilePath == "" {
		panic("Configuration file name is empty")
	}
	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration file. Error: %s\n", err))
	}
	var config LauncherConfig
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse configuration file. Error: %s\n", err))
	}

	finalConfig := generateConfig(&config)
	err = finalConfig.printConfig()
	if err != nil {
		fmt.Printf("Failed to write configuration: %s\n", err)
		os.Exit(100)
	}
}

func generateConfig(input *LauncherConfig) LauncherConfig {
	result := LauncherConfig{
		Linux:   input.Common,
		Mac:     input.Common,
		Windows: input.Common,
	}
	result.Linux.copyFrom(input.Linux)
	result.Mac.copyFrom(input.Mac)
	result.Windows.copyFrom(input.Windows)
	return result
}
