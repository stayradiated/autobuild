package main

import (
	"os"
	"os/exec"
)

type Git struct {
	URL    string
	Branch string
}

func (g *Git) Clone(dirname, filename string) error {
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", g.URL, filename)
	cmd.Dir = dirname
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
