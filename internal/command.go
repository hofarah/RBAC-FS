package internal

import (
	"github.com/hofarah/RBAC-FS/utils"
	"strings"
)

type Command struct {
	c    string
	args []string
}

const (
	ListCMD   = "list"
	OpenCMD   = "open"
	BackCMD   = "back"
	ReadCMD   = "read"
	ExitCMD   = "exit"
	WhoAmICmd = "whoami"

	ExecCMD         = "exec"
	ExecUsageString = "wrong COMMAND. usage: " + ExecCMD + " <file name>"

	CreateFileCMD         = "create-file"
	CreateFileUsageString = "wrong COMMAND. usage: " + CreateFileCMD + " <new file name>"

	CreateDirCMD         = "create-folder"
	CreateDirUsageString = "wrong COMMAND. usage: " + CreateDirCMD + " <new directory name>"

	RemoveFileCMD         = "remove-file"
	RemoveFileUsageString = "wrong COMMAND. usage: " + RemoveFileCMD + " <file name>"

	RemoveDirCMD         = "remove-folder"
	RemoveDirUsageString = "wrong COMMAND. usage: " + RemoveDirCMD + " <directory name>"

	HelpCMD = "help"

	AddUserCMD         = "add-user"
	AddUserUsageString = "wrong COMMAND. usage: " + AddUserCMD + " <user name> + Optional[Role]"

	RemoveUserCMD         = "remove-user"
	RemoveUserUsageString = "wrong COMMAND. usage: " + RemoveUserCMD + " <user name>"

	AddUserRoleCMD         = "add-user-role"
	AddUserRoleUsageString = "wrong COMMAND. usage: " + AddUserRoleCMD + " <user name>  <role name>"

	RemoveUserRoleCMD         = "remove-user-role"
	RemoveUserRoleUsageString = "wrong COMMAND. usage: " + RemoveUserRoleCMD + " <user name>  <role name>"

	AddRoleCMD         = "add-role"
	AddRoleUsageString = "wrong COMMAND. usage: " + AddRoleCMD + " <role name>"

	RemoveRoleCMD         = "remove-role"
	RemoveRoleUsageString = "wrong COMMAND. usage: " + RemoveRoleCMD + " <role name>"

	AddRoleForFileCMD            = "add-file-role"
	AddRoleForFileCMDUsageString = "wrong COMMAND. usage: " + AddRoleForFileCMD + " <file name> <role-name> <access>"

	RemoveRoleForFileCMD            = "remove-file-role"
	RemoveRoleForFileCMDUsageString = "wrong COMMAND. usage: " + RemoveRoleForFileCMD + " <file name> <role-name>"

	AddLabelForFileCMD            = "add-file-label"
	AddLabelForFileCMDUsageString = "wrong COMMAND. usage: " + AddLabelForFileCMD + " <file name> <label-name> <access>"

	AddLabelForUserCMD            = "add-user-label"
	AddLabelForUserCMDUsageString = "wrong COMMAND. usage: " + AddLabelForUserCMD + " <label name> <user-name>"

	//RemoveLabelForUserCMD            = "remove-user-label"
	//RemoveLabelForUserCMDUsageString = "wrong COMMAND. usage: " + RemoveLabelForUserCMD + " <label name> <user-name>"

	RemoveLabelCMD         = "remove-label"
	RemoveLabelUsageString = "wrong COMMAND. usage: " + RemoveLabelCMD + " <label name>"

	OpenUsageString = "wrong COMMAND. usage: " + OpenCMD + " <folder name>"
	ReadUsageString = "wrong COMMAND. usage: " + ReadCMD + " <file name>"
	HelpString      = "COMMANDS\n" + ListCMD + " - list of current directory\n" +
		BackCMD + " - go to parent directory\n" + OpenCMD + " - open folder\n" +
		ReadCMD + " - open file\n" +
		ExecCMD + " - execute a file\n" +
		WhoAmICmd + " - what is my username\n" +
		ExitCMD + " - exit user\n" +
		CreateDirCMD + " - create new directory\n" +
		CreateFileCMD + " - create new file\n" +
		RemoveDirCMD + " - remove directory\n" +
		RemoveFileCMD + " - remove file"

	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

var validCommands = []string{ListCMD, OpenCMD, BackCMD, ReadCMD, AddUserCMD, RemoveUserCMD, AddRoleCMD, RemoveRoleCMD, AddUserRoleCMD, RemoveUserRoleCMD, AddRoleForFileCMD, RemoveRoleForFileCMD, CreateDirCMD, CreateFileCMD, RemoveDirCMD, RemoveFileCMD, ExecCMD, ExitCMD, WhoAmICmd, AddLabelForFileCMD, RemoveLabelCMD, AddLabelForUserCMD, HelpCMD}

func (c *Command) Validate() bool {
	return utils.Contain(validCommands, c.c)
}

func (c *Command) GetType() string {
	return c.c
}
func (c *Command) IsEmpty() bool {
	return c.c == ""
}
func NewCommand(line string) Command {
	line = strings.TrimSpace(line)
	args := strings.Split(line, " ")
	return Command{c: args[0], args: args[1:]}
}
