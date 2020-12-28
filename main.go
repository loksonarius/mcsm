package main

import (
	"github.com/loksonarius/mcsm/cmd"
)

// I really wish this variable could've just been set under the cmd pkg >:V
var version = "unset"
var commit = "unset"

func main() {
	cmd.Execute(version, commit)
}
