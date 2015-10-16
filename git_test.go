package main

import (
	"fmt"
	"os"
	"testing"
)

var remote = "https://github.com/stayradiated/autobuild"
var branch = "master"

func TestGitClone(t *testing.T) {
	git := &Git{
		Remote: remote,
		Branch: branch,
	}

	os.MkdirAll("./builds", 0755)

	if err := git.Clone("./builds/_autobuild"); err != nil {
		fmt.Println(err)
	}
}

func TestGitStatus(t *testing.T) {
	git := &Git{remote, branch}
	if err := git.Status("./builds/_autobuild"); err != nil {
		panic(err)
	}
}

func TestGitCheckout(t *testing.T) {
	git := &Git{remote, branch}
	if err := git.Checkout("./builds/_autobuild"); err != nil {
		panic(err)
	}
}

func TestGitPull(t *testing.T) {
	git := &Git{remote, branch}
	if err := git.Pull("./builds/_autobuild"); err != nil {
		panic(err)
	}
}

func TestGitCurrentSHA(t *testing.T) {
	git := &Git{remote, branch}
	SHA, err := git.CurrentSHA("./builds/_autobuild")
	if err != nil {
		panic(err)
	}
	fmt.Println(SHA)
}
