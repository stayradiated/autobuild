package main

import "testing"

func TestGitClone(t *testing.T) {
	git := &Git{
		URL:    "https://github.com/stayradiated/autobuild",
		Branch: "master",
	}
	if err := git.Clone("./builds", "_autobuild"); err != nil {
		panic(err)
	}
}
