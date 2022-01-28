package internal

import "fmt"

type Printable interface {
	print()
}

type OPrint struct {
	text  string
	color string
}

func NewPrintable(text string, option ...string) Printable {
	color := colorReset
	if len(option) != 0 {
		color = option[0]
	}
	return &OPrint{text: text, color: color}
}
func NewError(text string) Printable {
	return &OPrint{text: text, color: colorRed}
}
func NewHelp(text string) Printable {
	return &OPrint{text: text, color: colorBlue}
}

func (c *OPrint) print() {
	fmt.Println(c.color, c.text, colorReset)
}
func TPrint(ptb Printable) {
	ptb.print()
}
