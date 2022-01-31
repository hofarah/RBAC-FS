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

func GetBuilder(builderType string) TBuilder {
	if builderType == "RBAC" {
		return &RBACTerminal{}
	}
	if builderType == "MAC" {
		return &MACTerminal{}
	}
	return nil
}
