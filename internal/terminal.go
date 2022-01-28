package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Terminal interface {
	HandleListCMD(args ...string) Printable
	HandleOpenCMD(arg string) Printable
	HandleReadCMD(arg string) Printable
	HandleBackCMD() Printable
	HandleAddUserCMD(args ...string) Printable
	getName() string
	getVersion() string
	setUser(*user)
}

type baseTerminal struct {
	name    string
	version string
	user    *user
}

func HandleCmd(terminal Terminal, cmd Command) {
	if cmd.IsEmpty() {
		return
	}
	if !cmd.Validate() {
		//TPrint(NewError("command is not valid - enter help to read instructions."))
		out, _ := exec.Command(cmd.c, cmd.args...).Output()
		TPrint(NewPrintable(string(out), OPrint{color: colorCyan}))
		return
	}
	switch cmd.GetType() {
	case HelpCMD:
		TPrint(NewHelp(HelpString))
		return
	case ListCMD:
		TPrint(terminal.HandleListCMD(cmd.args...))
		return
	case BackCMD:
		TPrint(terminal.HandleBackCMD())
		return
	case OpenCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(OpenUsageString))
			return
		}
		TPrint(terminal.HandleOpenCMD(cmd.args[0]))
		return
	case ReadCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(ReadUsageString))
			return
		}
		TPrint(terminal.HandleReadCMD(cmd.args[0]))
		return
	case AddUserCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(AddUserUsageString))
			return
		}
		TPrint(terminal.HandleAddUserCMD(cmd.args...))
		return
	}
}
func Listen(terminal Terminal) {
	InitialChecks(terminal)
	fmt.Println("starting terminal", terminal.getName(), "version:", terminal.getVersion())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(colorYellow, ">>>>> ", colorReset)
		reader.Reset(os.Stdin)
		line, _ := reader.ReadString('\n')
		HandleCmd(terminal, NewCommand(line[:len(line)-1]))
	}
}

func InitialChecks(terminal Terminal) *user {
	for {
		reader := bufio.NewReader(os.Stdin)
		for {
			TPrint(NewPrintable("please enter username:", OPrint{keepCurrentLine: true}))
			reader.Reset(os.Stdin)
			username := format(reader.ReadString('\n'))
			TPrint(NewPrintable("please enter password:", OPrint{keepCurrentLine: true}))
			reader.Reset(os.Stdin)
			password := format(reader.ReadString('\n'))
			u := GetUser(username, password)
			if u != nil {
				if u.username == "admin" && password == "admin" {
					TPrint(NewPrintable("you successfully logged in ,please reset your password:", OPrint{keepCurrentLine: true}))
					reader.Reset(os.Stdin)
					newPassword := format(reader.ReadString('\n'))
					u.UpdateUserPass(newPassword)
				}
				terminal.setUser(u)
				return u
			}
			TPrint(NewError("wrong input please try again"))
		}
	}
}
func format(str string, _ error) string {
	return str[:len(str)-1]
}
