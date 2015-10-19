package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"net/http"
	"regexp"
)

var githubRegexp = regexp.MustCompile(`refs\/heads\/(.+)`)

type Github struct {
	Ref string `json:"ref"`
}

func NewGithub(req *http.Request, secret string) (*Github, error) {
	event := new(Github)

	if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
		return event, err
	}

	return event, nil
}

func (g *Github) Branch() string {
	if match := githubRegexp.FindStringSubmatch(g.Ref); len(match) > 1 {
		return match[1]
	}
	return ""
}

func ParseWebhook(req *http.Request, secret string) (string, error) {
	event, err := NewGithub(req, secret)
	if err != nil {
		return "", err
	}
	return event.Branch(), nil
}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
