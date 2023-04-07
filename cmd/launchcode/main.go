//go:generate goversioninfo -icon=icons/launchcode.ico
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bitshifted/launchcode/config"
)

func main() {

	// get current application directory
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	appName := filepath.Base(exePath)
	//initialize logging to file
	logFileName := fmt.Sprintf("%s-*.log", appName)
	logFile, err := os.CreateTemp("", logFileName)
	if err != nil {
		fmt.Println("Failed to initialize log file")
	} else {
		log.SetOutput(logFile)
	}

	appDir := config.GetAppBaseDirectory(exePath)
	log.Printf("Application directory: %s\n", appDir)
	os.Chdir(appDir)

	jvmPath, err := config.FindJvmCommand(appDir)
	if err != nil {
		log.Println("Could not find java command")
	}

	args := config.GetCmdLineOptions()
	log.Printf("Command line: %v\n", args)

	binary := exec.Command(jvmPath, args...)

	out, execErr := binary.CombinedOutput()
	if execErr != nil {
		log.Printf("Error running Java process: %s\n", execErr.Error())
	}
	fmt.Println(string(out))
}
