//go:generate goversioninfo -icon=icons/launchcode.ico
package main

import (
	"fmt"
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
		fmt.Println("Could not find java command")
	}

	args := config.GetCmdLineOptions()
	fmt.Printf("Command line: %v\n", args)

	binary := exec.Command(jvmPath, args...)

	out, execErr := binary.CombinedOutput()
	if execErr != nil {
		fmt.Println(execErr)
	}
	fmt.Println(string(out))
}
