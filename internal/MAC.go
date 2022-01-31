package internal

import (
	"bufio"
	"context"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type MACTerminal struct {
	baseTerminal
}

func (r *MACTerminal) getName() string {
	return r.name
}
func (r *MACTerminal) getPath() string {
	return r.currentPath
}
func (r *MACTerminal) setUser(u *user) {
	r.user = u
}
func (r *MACTerminal) initPath() {
	p, _ := exec.Command("pwd").Output()
	r.currentPath = format(string(p), nil)
}
func (r *MACTerminal) setPath(path string) {
	r.currentPath = path
}
func (r *MACTerminal) addPath(p string) {
	r.currentPath = filepath.Join(r.currentPath, p)
}
func (r *MACTerminal) getVersion() string {
	return r.version
}
func (r *MACTerminal) setName() {
	r.name = "MAC"
}
func (r *MACTerminal) setVersion() {
	r.version = "1.0"
}
func (r *MACTerminal) HandleListCMD(args ...string) Printable {

	var response string
	dirs, _ := os.ReadDir(r.currentPath)
	for _, entitiy := range dirs {
		if MACAccess(r.user.id, 1, filepath.Join(r.currentPath, entitiy.Name())) || r.user.isAdmin() {
			response += entitiy.Name() + "\n"
		}
	}

	return NewPrintable(response)
}
func (r *MACTerminal) HandleOpenCMD(arg string) Printable {
	newPath := filepath.Join(r.currentPath, arg)
	_, err := os.ReadDir(newPath)
	if err != nil {
		return NewError(err.Error())
	}
	r.currentPath = newPath
	return NewPrintable("")
}
func (r *MACTerminal) HandleReadCMD(arg string) Printable {
	if !MACAccess(r.user.id, Read, filepath.Join(r.currentPath, arg)) || r.user.isAdmin() {
		return NewError("you don't have access to read this file")
	}
	out, err := os.ReadFile(filepath.Join(r.currentPath, arg))
	if err != nil {
		return NewError(err.Error())
	}
	return NewPrintable(string(out))
}
func (r *MACTerminal) HandleBackCMD() Printable {
	paths := strings.Split(r.currentPath, "/")
	//if len(paths)==0{}
	newPath := strings.Join(paths[:len(paths)-1], "/")
	_, err := os.ReadDir(newPath)
	if err != nil {
		return NewError(err.Error())
	}
	r.currentPath = newPath
	return NewPrintable("")
}
func (r *MACTerminal) HandleAddUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		role     string
		username string
	)
	username = args[0]
	if len(args) == 2 {
		role = args[1]
	} else {
		role = GetDefault("role")
	}
	TPrint(NewPrintable("please enter password of new user:", OPrint{keepCurrentLine: true}))
	reader := bufio.NewReader(os.Stdin)
	newPassword := format(reader.ReadString('\n'))
	err := NewUser(username, newPassword, role)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("new user successfully created")
}
func (r *MACTerminal) HandleRemoveUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	err := RemoveUser(args[0])
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("user successfully deleted")
}
func (r *MACTerminal) HandleAddRoleForUserCMD(args ...string) Printable {
	return NewError("invalid command")
}
func (r *MACTerminal) HandleRemoveRoleForUserCMD(args ...string) Printable {
	return NewError("invalid command")
}

func (r *MACTerminal) HandleAddRoleCMD(args ...string) Printable {
	return NewError("invalid command")
}
func (r *MACTerminal) HandleRemoveRoleCMD(args ...string) Printable {
	return NewError("invalid command")
}
func (r *MACTerminal) HandleAddRoleForFileCMD(args ...string) Printable {
	return NewError("invalid command")
}
func (r *MACTerminal) HandleRemoveRoleForFileCMD(args ...string) Printable {
	return NewError("invalid command")
}

func (r *MACTerminal) HandleCreateFileCMD(arg string) Printable {
	newFilePath := filepath.Join(r.currentPath, arg)
	_, err := os.ReadFile(newFilePath)
	if err == nil {
		return NewError("file already exist")
	}

	_, err = os.Create(newFilePath)
	if err != nil {
		return NewError("create file failed")
	}
	err = NewUserAccess(r.user.id, newFilePath)
	if err != nil {
		return NewError("create file failed")
	}
	return NewPrintable("file successfully created")
}
func (r *MACTerminal) HandleCreateDirCMD(arg string) Printable {
	newFilePath := filepath.Join(r.currentPath, arg)
	_, err := syscall.Open(newFilePath, syscall.SYS_READ, 0)
	if err == nil {
		return NewError("folder already exist")
	}

	err = os.Mkdir(newFilePath, os.ModePerm)
	if err != nil {
		return NewError("create folder failed")
	}
	err = NewUserAccess(r.user.id, newFilePath)
	if err != nil {
		return NewError("create folder failed")
	}
	return NewPrintable("folder successfully created")
}
func (r *MACTerminal) HandleRemoveFileCMD(arg string) Printable {
	filePathToDelete := filepath.Join(r.currentPath, arg)
	_, err := syscall.Open(filePathToDelete, syscall.SYS_READ, 0)
	if err != nil {
		return NewError("file does not exist to execute")
	}
	if !MACAccess(r.user.id, Write, filePathToDelete) && !r.user.isAdmin() {
		return NewError("you don't have Access to this file")
	}
	removeRecursively(r.user.id, filePathToDelete)

	err = os.Remove(filePathToDelete)
	if err != nil {
		return NewError("remove file failed")
	}
	return NewPrintable("file successfully removed")
}
func (r *MACTerminal) HandleExecCMD(arg string) Printable {

	filePathToExec := filepath.Join(r.currentPath, arg)
	_, err := syscall.Open(filePathToExec, syscall.SYS_READ, 0)
	if err != nil {
		return NewError("file does not exist to execute")
	}
	if !MACAccess(r.user.id, Execute, filePathToExec) && !r.user.isAdmin() {
		return NewError("you don't have Access to execute this file")
	}
	out, err := exec.Command(filePathToExec).Output()
	if err != nil {
		return NewError(err.Error())
	}
	return NewPrintable(string(out))
}
func (r *MACTerminal) HandleExitCMD() Printable {
	TPrint(NewPrintable("exited."))
	zap.L().Info("user exited", zap.Any("user", *r.user))
	InitialChecks(r)

	return NewPrintable("")
}
func (r *MACTerminal) HandleWhoAmICMD() Printable {
	return NewPrintable(r.user.username)
}
func (r *MACTerminal) HandleRemoveDirCMD(arg string) Printable {
	filePathToDelete := filepath.Join(r.currentPath, arg)

	if !MACAccess(r.user.id, Write, filePathToDelete) && !r.user.isAdmin() {
		return NewError("you don't have Access to this folder")
	}
	_, err := os.ReadDir(filePathToDelete)
	if err != nil {
		return NewError("folder does not exist")
	}
	removeRecursively(r.user.id, filePathToDelete)

	err = os.RemoveAll(filePathToDelete)
	if err != nil {
		return NewError("remove folder failed")
	}

	return NewPrintable("folder successfully removed")
}
func (r *MACTerminal) getTerminal() Terminal {
	return r
}

func (r *MACTerminal) HandleRemoveLabelCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	err := RemoveLabel(args[0])
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("label successfully deleted")
}
func (r *MACTerminal) HandleAddLabelForFileCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		filePath string
	)
	filePath = filepath.Join(r.currentPath, args[0])
	_, err := syscall.Open(filePath, syscall.SYS_READ, 0)
	if err != nil {
		return NewError("file does not exist")
	}

	err = NewLabel(args[1], filePath, AclToInt(args[2]), r.user.id)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("new label for file successfully created")
}
func (r *MACTerminal) HandleAddLabelForUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	l := GetLabel(args[0])
	if l == nil {
		return NewError("label not found")
	}
	u := GetUsername(args[1])
	if u == nil {
		return NewError("user not found")
	}
	err := NewUserLabel(u.id, l.id)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("new label for file successfully created")
}

func MACAccess(userID, level int, path string) bool {
	if UserAccess(userID, path) {
		return true
	} //it's wrong,but in project said this.
	if macAccess(userID, level, path) {
		return true
	}
	return false
}
func macAccess(userID, level int, path string) bool {
	var acl int
	err := conn.QueryRowContext(context.Background(), "select `level` from labels inner join labelUsers on (labelID=id) where userID=? AND access=?",
		userID, path).Scan(&acl)
	if err != nil {
		return false
	}
	return acl == level
}
