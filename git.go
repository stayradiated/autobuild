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

	cmd := Command{
		Command:      "git",
		Args:         []string{"clone", "--depth", "1", g.Remote, filename},
		HandleStdout: pipeToConsole,
	}
	return cmd.Exec(dirname)
}

func (g *Git) Status(repodir string) error {
	cmd := Command{
		Command:      "git",
		Args:         []string{"status"},
		HandleStdout: pipeToConsole,
	}
	return cmd.Exec(repodir)
}

func (g *Git) Fetch(repodir string) error {
	cmd := Command{
		Command:      "git",
		Args:         []string{"fetch"},
		HandleStdout: pipeToConsole,
	}
	return cmd.Exec(repodir)
}

func (g *Git) Checkout(repodir string) error {
	cmd := Command{
		Command:      "git",
		Args:         []string{"checkout", g.Branch},
		HandleStdout: pipeToConsole,
	}
	return cmd.Exec(repodir)
}

func (g *Git) Pull(repodir string) error {
	cmd := Command{
		Command:      "git",
		Args:         []string{"pull"},
		HandleStdout: pipeToConsole,
	}
	return cmd.Exec(repodir)
}

func (g *Git) CurrentSHA(repodir string) (string, error) {
	var SHA string

	getSHA := func(r io.Reader) {
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		SHA = strings.Split(scanner.Text(), " ")[1]
	}

	cmd := Command{
		Command:      "git",
		Args:         []string{"log", "-1", g.Branch},
		HandleStdout: getSHA,
	}

	return SHA, cmd.Exec(repodir)
}
