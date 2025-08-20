package version

var (
	Version = "dev"
	Commit  = ""
	Date    = ""
)

func Info() string {
	v := Version
	if Commit != "" {
		v += " (" + Commit + ")"
	}
	if Date != "" {
		v += " built on " + Date
	}
	return v
}
