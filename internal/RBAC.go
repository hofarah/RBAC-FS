package internal

import (
	"bufio"
	"os"
	"os/exec"
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
	//todo check permission of user
	cmd := exec.Command("ls", args...)
	out, _ := cmd.Output()
	//todo check output

	/////////////fixme OR

	//dirs,_:=os.ReadDir(args[0])
	//for _,dir:=range dirs{
	//	//todo check permission and delete
	//}

	return NewPrintable(string(out))
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
func (r *RBACTerminal) getTerminal() Terminal {
	return r
}
