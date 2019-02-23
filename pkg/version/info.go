// Package version contains build information.
package version

import "runtime"

// Build information. Populated at build-time.
var (
	version   = "dev"
	revision  string
	buildUser string
	buildDate string
	goVersion = runtime.Version()
)

// Information holds app version info
type Information struct {
	Version   string `json:"version"`
	Revision  string `json:"revision"`
	Branch    string `json:"branch"`
	BuildUser string `json:"build_user"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
}

// Info returns app version info
func Info() Information {
	return Information{
		Version:   version,
		Revision:  revision,
		BuildUser: buildUser,
		BuildDate: buildDate,
		GoVersion: goVersion,
	}
}

// String return version information as string
func (i Information) String() string {
	return "Version: " + i.Version +
		", Revision: " + i.Revision +
		", BuildUser: " + i.BuildUser +
		", BuildDate: " + i.BuildDate +
		", GoVersion: " + i.GoVersion
}
