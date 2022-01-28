package main

import (
	"github.com/hofarah/RBAC-FS/internal"
)

func main() {
	internal.Listen(internal.BuildTerminal(internal.GetBuilder("RBAC")))
}
