package internal

type RBACTerminal struct {
	baseTerminal
}

func (r *RBACTerminal) getName() string {
	return r.name
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
func (r *RBACTerminal) HandleListCMD() Printable {
	//todo implement me
	return NewPrintable("list of current directory")
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
func (r *RBACTerminal) getTerminal() Terminal {
	return r
}
