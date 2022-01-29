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
	ListCMD = "list"
	OpenCMD = "open"
	BackCMD = "back"
	ReadCMD = "read"

	CreateFileCMD         = "create-file" //todo
	CreateFileUsageString = "wrong COMMAND. usage: " + CreateFileCMD + " <new file name>"

	CreateDirCMD         = "create-folder" //todo
	CreateDirUsageString = "wrong COMMAND. usage: " + CreateDirCMD + " <new folder name>"

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
	AddRoleForFileCMDUsageString = "wrong COMMAND. usage: " + AddRoleForFileCMD + " <file name> <role-name>"

	RemoveRoleForFileCMD            = "remove-file-role"
	RemoveRoleForFileCMDUsageString = "wrong COMMAND. usage: " + RemoveRoleForFileCMD + " <file name> <role-name>"

	OpenUsageString = "wrong COMMAND. usage: " + OpenCMD + " <folder name>"
	ReadUsageString = "wrong COMMAND. usage: " + ReadCMD + " <file name>"
	HelpString      = "COMMANDS\n" + ListCMD + " - list of current directory\n" + BackCMD + " - go to parent directory\n" + OpenCMD + " - open folder\n" + ReadCMD + " - open file"
	colorReset      = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

var validCommands = []string{ListCMD, OpenCMD, BackCMD, ReadCMD, AddUserCMD, RemoveUserCMD, AddRoleCMD, RemoveRoleCMD, AddUserRoleCMD, RemoveUserRoleCMD, AddRoleForFileCMD, RemoveRoleForFileCMD, CreateDirCMD, CreateFileCMD, HelpCMD}

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
