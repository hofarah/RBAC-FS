package internal

import "fmt"

type Printable interface {
	print()
}

type OPrint struct {
	text            string
	color           string
	keepCurrentLine bool
}

func NewPrintable(text string, option ...OPrint) Printable {
	color := colorReset
	keepCurrentLine := false
	if len(option) != 0 {
		color = option[0].color
		keepCurrentLine = option[0].keepCurrentLine
	}
	return &OPrint{text: text, color: color, keepCurrentLine: keepCurrentLine}
}
func NewError(text string) Printable {
	return &OPrint{text: text, color: colorRed}
}
func NewHelp(text string) Printable {
	return &OPrint{text: text, color: colorBlue}
}

func (c *OPrint) print() {
	if c.keepCurrentLine {
		fmt.Print(c.color, c.text, colorReset)
	} else {
		fmt.Println(c.color, c.text, colorReset)
	}
}
func TPrint(ptb Printable) {
	ptb.print()
}
