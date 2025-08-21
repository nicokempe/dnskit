package version

var (
	// Version is the semantic version of the DNSKit build.
	Version = "dev"
	// Commit is the short git commit hash for the build.
	Commit = ""
	// Date is the build date in UTC.
	Date = ""
)

// Info returns a human-readable version string, including commit and build date when available.
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
