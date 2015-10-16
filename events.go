package main

import (
	"encoding/json"
	"io"
	"regexp"
)

var githubRegexp = regexp.MustCompile(`refs\/heads\/(.+)`)

type GithubPush struct {
	Ref string `json:"ref"`
}

func (g *GithubPush) Branch() string {
	if match := githubRegexp.FindStringSubmatch(g.Ref); len(match) > 1 {
		return match[1]
	}
	return ""
}

func ParseWebhook(r io.Reader) (string, error) {
	event := new(GithubPush)

	if err := json.NewDecoder(r).Decode(&event); err != nil {
		return "", err
	}

	return event.Branch(), nil
}
