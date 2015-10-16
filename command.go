package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type HandlePipe func(io.Reader)

func pipeToConsole(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("| %s\n", scanner.Text())
	}
}

type Command struct {
	Command      string
	Args         []string
	HandleStdout HandlePipe
}

func (c *Command) Exec(dirname string) error {
	cmd := exec.Command(c.Command, c.Args...)
	cmd.Dir = dirname

	fmt.Printf("> %s %s\n", c.Command, strings.Join(c.Args, " "))

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}
	go c.HandleStdout(cmdReader)

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}

	return nil
}
