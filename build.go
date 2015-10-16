package main

import (
	"bytes"
	"path"
	"strings"
	"text/template"
)

type Build struct {
	Dir     string
	Command string
	Args    string
}

type BuildVariables struct {
	Version string
}

func (b *Build) Exec(dirname string, vars *BuildVariables) error {
	tmpl, err := template.New("args").Parse(b.Args)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, vars)
	if err != nil {
		panic(err)
	}

	cmd := Command{
		Command:      b.Command,
		Args:         strings.Split(buf.String(), " "),
		HandleStdout: pipeToConsole,
	}

	return cmd.Exec(path.Join(dirname, b.Dir))
}
