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
	HandleRemoveUserCMD(args ...string) Printable
	HandleAddRoleCMD(args ...string) Printable
	HandleRemoveRoleCMD(args ...string) Printable
	HandleAddRoleForFileCMD(args ...string) Printable
	HandleRemoveRoleForFileCMD(args ...string) Printable
	HandleAddRoleForUserCMD(args ...string) Printable
	HandleRemoveRoleForUserCMD(args ...string) Printable
	getName() string
	getPath() string
	getVersion() string
	setUser(*user)
	initPath()
	addPath(path string)
	setPath(path string)
}

type baseTerminal struct {
	name        string
	version     string
	user        *user
	currentPath string
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
	case RemoveUserCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(RemoveUserUsageString))
			return
		}
		TPrint(terminal.HandleRemoveUserCMD(cmd.args...))
		return
	case AddRoleCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(AddRoleUsageString))
			return
		}
		TPrint(terminal.HandleAddRoleCMD(cmd.args...))
		return
	case RemoveRoleCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(RemoveRoleUsageString))
			return
		}
		TPrint(terminal.HandleRemoveRoleCMD(cmd.args...))
		return
	case AddRoleForFileCMD:
		if len(cmd.args) <= 1 {
			TPrint(NewHelp(AddRoleForFileCMDUsageString))
			return
		}
		TPrint(terminal.HandleAddRoleForFileCMD(cmd.args...))
		return
	case RemoveRoleForFileCMD:
		if len(cmd.args) <= 1 {
			TPrint(NewHelp(RemoveRoleForFileCMDUsageString))
			return
		}
		TPrint(terminal.HandleRemoveRoleForFileCMD(cmd.args...))
		return
	case AddUserRoleCMD:
		if len(cmd.args) <= 1 {
			TPrint(NewHelp(AddUserRoleUsageString))
			return
		}
		TPrint(terminal.HandleAddRoleForUserCMD(cmd.args...))
		return
	case RemoveUserRoleCMD:
		if len(cmd.args) <= 1 {
			TPrint(NewHelp(RemoveUserRoleUsageString))
			return
		}
		TPrint(terminal.HandleRemoveRoleForUserCMD(cmd.args...))
		return
	}
}
func Listen(terminal Terminal) {
	InitialChecks(terminal)
	fmt.Println("starting terminal", terminal.getName(), "version:", terminal.getVersion())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(colorYellow, terminal.getPath(), ">> ", colorReset)
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
				terminal.initPath()
				return u
			}
			TPrint(NewError("wrong input please try again"))
		}
	}
}
func format(str string, _ error) string {
	return str[:len(str)-1]
}
