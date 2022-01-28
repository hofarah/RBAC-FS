package internal

type Terminal interface {
	HandleCmd(command) string
}

type baseTerminal struct {
	name    string
	version string
}
