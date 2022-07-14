package config

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const DEFAULT_JAVA_DIR = "jre"
const DEFAULT_BIN_DIR = "bin"

var (
	//go:embed embed/jvmopts.txt
	jvmOptions string
	//go:embed embed/jvmprops.txt
	jvmProperties string
	//go:embed embed/splash.txt
	splashScreen string
	//go:embed embed/modulepath.txt
	modulePath string
	//go:embed embed/addmodules.txt
	addModulesData string
	//go:embed embed/module.txt
	module string
	//go:embed embed/classpath.txt
	classpath string
	//go:embed embed/mainclass.txt
	mainClass string
	//go:embed embed/jar.txt
	jar string
	//go:embed embed/args.txt
	arguments string
)

func FindJvmCommand(appDir string) (string, error) {
	jvmDir := calculateTargetJvmDir(appDir)
	fmt.Println(jvmDir)
	// attempt to find java executable
	javaExecPath := findJavaExecutable(jvmDir)
	fmt.Printf("java exec path: %s\n", javaExecPath)

	return javaExecPath, nil

}

func calculateTargetJvmDir(appDir string) string {
	return filepath.Join(appDir, DEFAULT_JAVA_DIR)

}

func findJavaExecutable(origin string) string {
	var execName = "java"
	if runtime.GOOS == "windows" {
		execName = "javaw"
	}
	return filepath.Join(origin, DEFAULT_BIN_DIR, execName)

}

func GetCmdLineOptions() []string {
	options := make([]string, 0)
	options = append(options, setJvmOptions()...)
	options = append(options, setJvmProperties()...)
	options = append(options, setSplashScreen()...)
	options = append(options, setModulePath()...)
	options = append(options, addModules()...)
	options = append(options, setModule()...)
	options = append(options, setClasspath()...)
	options = append(options, setMainClass()...)
	options = append(options, setJar()...)
	options = append(options, setArguments()...)

	return options
}

func setModulePath() []string {
	if len(modulePath) > 0 {
		return []string{"--module-path", strings.TrimSpace(modulePath)}
	}
	return []string{}

}

func setModule() []string {
	if len(module) > 0 && len(mainClass) > 0 {
		var builder strings.Builder
		builder.WriteString(strings.TrimSpace(module))
		builder.WriteString("/")
		builder.WriteString(strings.TrimSpace(mainClass))
		return []string{"--module", builder.String()}
	}
	return []string{}

}

func addModules() []string {
	if len(addModulesData) > 0 {
		return []string{fmt.Sprintf("--add-modules=%s", strings.TrimSpace(addModulesData))}
	}
	return []string{}
}

func setClasspath() []string {
	if len(classpath) > 0 {
		return []string{"-classpath", strings.TrimSpace(classpath)}
	}
	return []string{}
}

func setJvmOptions() []string {
	return strings.Fields(jvmOptions)
}

func setJvmProperties() []string {
	return strings.Fields(jvmProperties)
}

func setArguments() []string {
	if len(arguments) > 0 {
		return strings.Fields(arguments)
	}
	return []string{}
}

func setMainClass() []string {
	if len(jar) > 0 || len(module) > 0 {
		return []string{}
	}
	return []string{strings.TrimSpace(mainClass)}
}

func setJar() []string {
	if len(jar) > 0 {
		return []string{"-jar", strings.TrimSpace(jar)}
	}
	return []string{}
}

func setSplashScreen() []string {
	if len(splashScreen) > 0 && strings.HasPrefix(splashScreen, "-splash:") {

		return []string{splashScreen}
	}
	return []string{}
}
