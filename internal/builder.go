package internal

type TBuilder interface {
	setName()
	setVersion()
	getTerminal() Terminal
}

func BuildTerminal(builder TBuilder) Terminal {
	builder.setVersion()
	builder.setName()
	return builder.getTerminal()
}

func getBuilder(builderType string) TBuilder {
	if builderType == "RBAC" {
		return &RBACTerminal{}
	}
	if builderType == "MAC" {
		//todo
	}
	return nil
}
