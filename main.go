//go:generate goversioninfo
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
	// httpClient := client.Init()
	application := config.LoadConfigFile()

	// status := client.CheckUpdateStatus(application, httpClient)
	// fmt.Println(status)
	jvmPath, err := config.FindJvmCommand(application, appDir)
	if(err != nil) {
		fmt.Println("Could not find java command")
	}
	fmt.Println(jvmPath)

	args := config.GetCmdLineOptions(application)
	fmt.Println(args)

	binary := exec.Command(jvmPath, args...)

	out, execErr := binary.CombinedOutput()
	if execErr != nil {
		fmt.Println(execErr)
	}
	fmt.Println(string(out))
}
