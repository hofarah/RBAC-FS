package main

import (
	"github.com/hofarah/RBAC-FS/internal"
	_ "github.com/hofarah/RBAC-FS/logger"
)

func main() {
	internal.Listen(internal.BuildTerminal(internal.GetBuilder("MAC")))
}
