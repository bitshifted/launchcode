package config

const (
	syncroMainClass = "co.bitshifted.xapps.syncro.Syncro"
)

func GetSyncroCmdOptions() []string {
	return []string{"-classpath", "cp/*", syncroMainClass}
}
