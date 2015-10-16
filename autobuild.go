package main

import (
	"os"
)

type AutoBuild struct {
	Git   *Git
	Build *Build
}

func (a *AutoBuild) Run(reponame string) error {
	exists, err := checkDirExists(reponame)
	if err != nil {
		return err
	}

	if !exists {
		if err := a.Git.Clone(reponame); err != nil {
			return err
		}
		if err := a.Git.Checkout(reponame); err != nil {
			return err
		}
	} else {
		if err := a.Git.Checkout(reponame); err != nil {
			return err
		}
		if err := a.Git.Pull(reponame); err != nil {
			return err
		}
	}

	SHA, err := a.Git.CurrentSHA(reponame)
	if err != nil {
		return err
	}

	if err := a.Build.Exec(reponame, &BuildVariables{
		Version: SHA[0:6],
	}); err != nil {
		return err
	}

	return nil
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
