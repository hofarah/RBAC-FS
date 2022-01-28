package internal

type command struct {
	c    string
	args []string
}

func (c *command) Do() string {
	//todo run command
	return ""
}

type Parser interface {
	Parse() string
}
