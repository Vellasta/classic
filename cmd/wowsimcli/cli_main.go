package main

import (
	"github.com/Vellasta/classic/cmd/wowsimcli/cmd"
	"github.com/Vellasta/classic/sim"
)

func init() {
	sim.RegisterAll()
}

// Version information.
// This variable is set by the makefile in the release process.
var Version string

func main() {
	cmd.Execute(Version)
}
