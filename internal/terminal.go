package internal

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Terminal interface {
	HandleListCMD(args ...string) Printable
	HandleOpenCMD(arg string) Printable
	HandleReadCMD(arg string) Printable
	HandleExecCMD(arg string) Printable
	HandleExitCMD() Printable
	HandleWhoAmICMD() Printable
	HandleBackCMD() Printable
	HandleAddUserCMD(args ...string) Printable
	HandleCreateFileCMD(args string) Printable
	HandleCreateDirCMD(args string) Printable
	HandleRemoveFileCMD(args string) Printable
	HandleRemoveDirCMD(args string) Printable
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
	zap.L().Info("new command", zap.Any("command type", cmd.GetType()), zap.Any("command args", cmd.args), zap.Any("terminal", terminal))
	if cmd.IsEmpty() {
		return
	}
	if !cmd.Validate() { //preventing command injection
		TPrint(NewError("invalid command"))
		TPrint(NewHelp(HelpString))
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
		if len(cmd.args) <= 2 {
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
	case CreateFileCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(CreateFileUsageString))
			return
		}
		TPrint(terminal.HandleCreateFileCMD(cmd.args[0]))
		return
	case CreateDirCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(CreateDirUsageString))
			return
		}
		TPrint(terminal.HandleCreateDirCMD(cmd.args[0]))
		return
	case RemoveFileCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(RemoveFileUsageString))
			return
		}
		TPrint(terminal.HandleRemoveFileCMD(cmd.args[0]))
		return
	case RemoveDirCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(RemoveDirUsageString))
			return
		}
		TPrint(terminal.HandleRemoveDirCMD(cmd.args[0]))
		return
	case ExecCMD:
		if len(cmd.args) == 0 {
			TPrint(NewHelp(ExecUsageString))
			return
		}
		TPrint(terminal.HandleExecCMD(cmd.args[0]))
		return
	case ExitCMD:
		TPrint(terminal.HandleExitCMD())
		return
	case WhoAmICmd:
		TPrint(terminal.HandleWhoAmICMD())
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
			u := checkPublicKey(username) //check entrance with public key
			if u != nil {
				terminal.setUser(u)
				terminal.initPath()
				zap.L().Info("new entrance", zap.String("type", "public-key"), zap.Any("user", u.username))
				TPrint(NewPrintable("login via public key", OPrint{color: colorCyan}))
				return u
			}
			TPrint(NewPrintable("please enter password:", OPrint{keepCurrentLine: true}))
			reader.Reset(os.Stdin)
			password := format(reader.ReadString('\n'))
			u = GetUser(username, password)
			if u != nil {
				if u.username == "admin" && password == "admin" {
					TPrint(NewPrintable("you successfully logged in ,please reset your password:", OPrint{keepCurrentLine: true}))
					reader.Reset(os.Stdin)
					newPassword := format(reader.ReadString('\n'))
					u.UpdateUserPass(newPassword)
				}
				terminal.setUser(u)
				terminal.initPath()
				zap.L().Info("new entrance", zap.String("type", "password"), zap.Any("user", u.username))

				return u
			}
			TPrint(NewError("wrong input please try again"))
		}
	}
}
func format(str string, _ error) string {
	return str[:len(str)-1]
}
func checkPublicKey(username string) *user {
	home, _ := os.UserHomeDir()
	pubKey, err := os.ReadFile(filepath.Join(home, "/.ssh/id_rsa.pub"))
	if err != nil {
		return nil
	}
	return GetUserWithPublicKey(username, string(pubKey))

}
