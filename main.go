package main

import (
	"fmt"

	"github.com/gesquive/shield/cmd"
)

var version = "v0.8.0-git"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("shield %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
