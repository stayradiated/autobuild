package main

import (
	"fmt"
	"github.com/stayradiated/slacker"
	"log"
	"net/http"
	"os"
)

type AutoBuild struct {
	Dir     string
	Git     *Git
	Build   *Build
	Webhook *Webhook
	Slacker *slacker.Slacker
}

func (a *AutoBuild) Run() error {
	if err := a.fetchAndBuild(); err != nil {
		a.Slacker.Send("Build failed! Check the error logs for more info.")
		return err
	}
	return nil
}

func (a *AutoBuild) fetchAndBuild() error {
	repo := a.Dir

	exists, err := checkDirExists(repo)
	if err != nil {
		return err
	}

	if !exists {
		if err := a.Git.Clone(repo); err != nil {
			return err
		}
		if err := a.Git.Checkout(repo); err != nil {
			return err
		}
	} else {
		if err := a.Git.Fetch(repo); err != nil {
			return err
		}
		if err := a.Git.Checkout(repo); err != nil {
			return err
		}
		if err := a.Git.Pull(repo); err != nil {
			return err
		}
		if err := a.Git.UpdateSubmodules(repo); err != nil {
			return err
		}
	}

	SHA, err := a.Git.CurrentSHA(repo)
	if err != nil {
		return err
	}
	SHA = SHA[0:6]

	a.Slacker.Send(fmt.Sprintf("Starting to build version %s", SHA))

	if err := a.Build.Exec(repo, &BuildVariables{
		Version: SHA,
	}); err != nil {
		return err
	}

	a.Slacker.Send(fmt.Sprintf("Finished building version %s", SHA))

	return nil
}

func (a *AutoBuild) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if err := a.Webhook.Handle(resp, req, a.HandlePush); err != nil {
		log.Println(err)
	}
}

func (a *AutoBuild) HandlePush(branch string) {
	if branch == a.Git.Branch {
		a.Run()
	}
}

// exists returns whether the given file or directory exists or not
func checkDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
