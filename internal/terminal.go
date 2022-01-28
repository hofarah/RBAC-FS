package internal

import (
	"bufio"
	"fmt"
	"os"
)

type Terminal interface {
	HandleListCMD() Printable
	HandleOpenCMD(arg string) Printable
	HandleReadCMD(arg string) Printable
	HandleBackCMD() Printable
	getName() string
	getVersion() string
}

type baseTerminal struct {
	name    string
	version string
}

func HandleCmd(terminal Terminal, cmd Command) {
	if !cmd.Validate() {
		TPrint(NewError("command is not valid - enter help to read instructions."))
		return
	}
	switch cmd.GetType() {
	case HelpCMD:
		TPrint(NewHelp(HelpString))
		return
	case ListCMD:
		terminal.HandleListCMD()
		return
	case BackCMD:
		terminal.HandleBackCMD()
		return
	case OpenCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(OpenUsageString))
			return
		}
		terminal.HandleOpenCMD(cmd.args[0])
		return
	case ReadCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(ReadUsageString))
			return
		}
		terminal.HandleReadCMD(cmd.args[0])
		return
	}
}
func Listen(terminal Terminal) {
	fmt.Println("starting terminal", terminal.getName(), "version:", terminal.getVersion())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(colorYellow, ">>>>> ", colorReset)
		reader.Reset(os.Stdin)
		line, _ := reader.ReadString('\n')
		HandleCmd(terminal, NewCommand(line[:len(line)-1]))
	}
}
