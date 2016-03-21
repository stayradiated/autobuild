package main

import (
	"bufio"
	"io"
	"path"
	"strings"
)

type Git struct {
	Remote string
	Branch string
}

func (g *Git) Clone(reponame string) error {
	filename := path.Base(reponame)
	dirname := path.Dir(reponame)

	cmd := NewCommand("git", "clone", g.Remote, filename)
	return cmd.Exec(dirname)
}

func (g *Git) Status(repodir string) error {
	cmd := NewCommand("git", "status")
	return cmd.Exec(repodir)
}

func (g *Git) Fetch(repodir string) error {
	cmd := NewCommand("git", "fetch")
	return cmd.Exec(repodir)
}

func (g *Git) Checkout(repodir string) error {
	cmd := NewCommand("git", "checkout", g.Branch)
	return cmd.Exec(repodir)
}

func (g *Git) Pull(repodir string) error {
	cmd := NewCommand("git", "pull")
	return cmd.Exec(repodir)
}

func (g *Git) UpdateSubmodules(repodir string) error {
	cmd := NewCommand("git", "submodule", "update", "--init", "--recursive")
	return cmd.Exec(repodir)
}

func (g *Git) CurrentSHA(repodir string) (string, error) {
	var SHA string

	getSHA := func(r io.Reader) {
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		SHA = strings.Split(scanner.Text(), " ")[1]
	}

	cmd := NewCommand("git", "log", "-1", g.Branch)
	cmd.HandleStdout = getSHA
	return SHA, cmd.Exec(repodir)
}
