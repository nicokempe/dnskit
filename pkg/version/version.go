package version

// Version represents the semantic version of the DNSKit build.
// Commit stores the short git commit hash used for the build.
// Date records the build date in UTC.
var (
	Version = "dev"
	Commit  = ""
	Date    = ""
)

// Info returns a human-readable version string including commit and build date
// information when available.
func Info() string {
	versionInfo := Version
	if Commit != "" {
		versionInfo += " (" + Commit + ")"
	}
	if Date != "" {
		versionInfo += " built on " + Date
	}
	return versionInfo
}
