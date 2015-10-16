package main

import (
	"testing"
)

func TestBuildExec(t *testing.T) {

	build := &Build{
		Dir:     ".",
		Command: "echo",
		Args:    "--version {{.Version}} --deploy",
	}

	build.Exec(&BuildVariables{
		Version: "1.2.3",
	})

}
