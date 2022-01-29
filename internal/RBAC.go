package internal

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type RBACTerminal struct {
	baseTerminal
}

func (r *RBACTerminal) getName() string {
	return r.name
}
func (r *RBACTerminal) setUser(u *user) {
	r.user = u
}
func (r *RBACTerminal) initPath() {
	p, _ := exec.Command("pwd").Output()
	r.currentPath = format(string(p), nil)
}
func (r *RBACTerminal) setPath(path string) {
	r.currentPath = path
}
func (r *RBACTerminal) addPath(p string) {
	r.currentPath = filepath.Join(r.currentPath, p)
}
func (r *RBACTerminal) getVersion() string {
	return r.version
}
func (r *RBACTerminal) setName() {
	r.name = "RBAC"
}
func (r *RBACTerminal) setVersion() {
	r.version = "1.0"
}
func (r *RBACTerminal) HandleListCMD(args ...string) Printable {

	var response string
	dirs, _ := os.ReadDir(r.currentPath)
	for _, entitiy := range dirs {
		if Access(r.user.id, filepath.Join(r.currentPath, entitiy.Name())) {
			response += entitiy.Name() + "\n"
		} else { //recursive check
			currentPath := r.currentPath
			for currentPath != "" {
				if Access(r.user.id, currentPath) {
					response += entitiy.Name() + "\n"
					break
				}
				paths := strings.Split(currentPath, "/")
				if len(paths) == 0 {
					break
				}
				currentPath = strings.Join(paths[:len(paths)-1], "/")
			}
		}
	}

	return NewPrintable(response)
}
func (r *RBACTerminal) HandleOpenCMD(arg string) Printable {
	//todo implement me
	return NewPrintable("opening folder:" + arg)
}
func (r *RBACTerminal) HandleReadCMD(arg string) Printable {
	//todo implement me
	return NewPrintable("reading file:" + arg)
}
func (r *RBACTerminal) HandleBackCMD() Printable {
	//todo implement me
	return NewPrintable("back to parent folder")
}
func (r *RBACTerminal) HandleAddUserCMD(args ...string) Printable {
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
func (r *RBACTerminal) HandleRemoveUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	err := RemoveUser(args[0])
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("user successfully deleted")
}
func (r *RBACTerminal) HandleAddRoleForUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		rol      string
		username string
	)
	username = args[0]
	rol = args[1]
	rl := GetRole(rol)
	if rl == nil {
		return NewError("role not found")
	}
	us := GetUsername(username)
	if us == nil {
		return NewError("user not found")
	}
	_, err := conn.ExecContext(context.Background(), "insert into userRoles (roleID,userID) values (?,?)", rl.id, us.id)
	if err != nil {
		return NewError("an error accuerd")
	}
	return NewPrintable("new user role successfully added")
}
func (r *RBACTerminal) HandleRemoveRoleForUserCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		rol      string
		username string
	)
	username = args[0]
	rol = args[1]
	rl := GetRole(rol)
	if rl == nil {
		return NewError("role not found")
	}
	us := GetUsername(username)
	if us == nil {
		return NewError("user not found")
	}
	_, err := conn.ExecContext(context.Background(), "delete from userRoles where roleID=? AND userID=? ", rl.id, us.id)
	if err != nil {
		return NewError("an error accuerd")
	}
	return NewPrintable("user role successfully deleted")
}

func (r *RBACTerminal) HandleAddRoleCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		rol string
	)
	rol = args[0]
	err := NewRole(rol)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("new role successfully created")
}
func (r *RBACTerminal) HandleRemoveRoleCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	err := RemoveRole(args[0])
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("role successfully deleted")
}
func (r *RBACTerminal) HandleAddRoleForFileCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		filePath string
		rol      string
	)
	filePath = filepath.Join(r.currentPath, args[0])
	rol = args[1]
	rr := GetRole(rol)
	if rr == nil {
		return NewError("role not found")
	}
	err := NewRoleAccess(rr.id, filePath)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("new role for file successfully created")
}
func (r *RBACTerminal) HandleRemoveRoleForFileCMD(args ...string) Printable {
	if !r.user.isAdmin() {
		return NewError("you are not admin!!")
	}
	var (
		filePath string
		rol      string
	)
	filePath = filepath.Join(r.currentPath, args[0])
	rol = args[1]
	rr := GetRole(rol)
	if rr == nil {
		return NewError("role not found")
	}
	err := RemoveRoleAccess(rr.id, filePath)
	if err != nil {
		return NewPrintable("an error accrued", OPrint{color: colorRed})
	}
	return NewPrintable("role for file successfully deleted")
}
func (r *RBACTerminal) getTerminal() Terminal {
	return r
}
