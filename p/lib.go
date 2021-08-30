package p

var (
	name    = "b"
	version = "v1.1.0"
)

func Version() string {
	return name + ": " + version
}
