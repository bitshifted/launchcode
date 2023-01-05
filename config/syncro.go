package config

import "fmt"

const (
	syncroMainClass = "co.bitshifted.appforge.syncro.Syncro"
)

func GetSyncroCmdOptions(launcherFilePath string) []string {
	options := make([]string, 0)
	options = append(options, setSplashScreen()...)
	options = append(options, setClasspath()...)
	options = append(options, []string{syncroMainClass, fmt.Sprintf("--launcher-file=%s", launcherFilePath)}...)
	return options
}
