package config

import "encoding/xml"

type Application struct {
	XMLName xml.Name `xml:"application"`
	Version string   `xml:"version,attr"`
	Release string   `xml:"release,attr"`

	Server Server `xml:"server"`
	Jvm    Jvm    `xml:"jvm"`
}

type Server struct {
	XMLName xml.Name `xml:"server"`
	BaseUrl string   `xml:"base-url,attr"`
}

type Jvm struct {
	XMLName       xml.Name `xml:"jvm"`
	JvmDir        string   `xml:"jvm-dir"`
	ModulePath    string   `xml:"module-path"`
	Module        string   `xml:"module"`
	AddModules    string   `xml:"add-modules"`
	Classpath     string   `xml:"classpath"`
	JvmOptions    string   `xml:"jvm-options"`
	JvmProperties string   `xml:"jvm-properties"`
	MainClass     string   `xml:"main-class"`
	Arguments     string   `xml:"args"`
	Jar           string   `xml:"jar"`
	SplashScreen  string   `xml:"splash-screen"`
}
