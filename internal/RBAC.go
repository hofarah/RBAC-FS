package internal

type RBACTerminal struct {
	baseTerminal
}

func (r *RBACTerminal) HandleCmd(command) string {
	//todo handle command
	return ""
}
func (r *RBACTerminal) setName() {
	r.name = "RBAC Terminal"
}
func (r *RBACTerminal) setVersion() {
	r.version = "1.0"
}
func (r *RBACTerminal) getTerminal() Terminal {
	return r
}
