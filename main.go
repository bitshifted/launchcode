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
	appDir := filepath.Dir(exePath)
	fmt.Println(appDir)
	os.Chdir(appDir)

	jvmPath, err := config.FindJvmCommand(appDir)
	if err != nil {
		log.Println("Could not find java command")
	}

	syncroArgs := config.GetSyncroCmdOptions()
	fmt.Printf("Syncro args: %v\n", syncroArgs)
	syncro := exec.Command(jvmPath, syncroArgs...)
	syncroOut, err := syncro.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run syncro: %s\n", err.Error())
	}
	log.Println(string(syncroOut))

	args := config.GetCmdLineOptions()
	log.Printf("Command line: %v\n", args)

	binary := exec.Command(jvmPath, args...)

	out, execErr := binary.CombinedOutput()
	if execErr != nil {
		log.Printf("Error running Java process: %s\n", execErr.Error())
	}
	log.Println(string(out))
}
