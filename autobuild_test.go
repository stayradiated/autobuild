package main

import "testing"

func TestAutoBuildRun(t *testing.T) {
	autobuild := &AutoBuild{
		Git: &Git{
			Remote: "https://github.com/stayradiated/autobuild",
			Branch: "master",
		},
		Build: &Build{
			Dir:     ".",
			Command: "echo",
			Args:    "--version {{.Version}} --deploy",
		},
	}
	autobuild.Run("./builds/autobuild")
}
