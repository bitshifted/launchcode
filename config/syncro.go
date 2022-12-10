package config

const (
	syncroMainClass = "co.bitshifted.appforge.syncro.Syncro"
)

func GetSyncroCmdOptions() []string {
	options := make([]string, 0)
	options = append(options, setSplashScreen()...)
	options = append(options, []string{"-classpath", "cp/*", syncroMainClass}...)
	return options
}
