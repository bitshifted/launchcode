package config

const (
	syncroMainClass = "co.bitshifted.appforge.syncro.Syncro"
)

func GetSyncroCmdOptions() []string {
	return []string{"-classpath", "cp/*", syncroMainClass}
}
