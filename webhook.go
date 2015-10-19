package main

import (
	"net/http"
)

type PushHandler func(branch string)

type Webhook struct {
	Secret string
}

func (w *Webhook) Handle(resp http.ResponseWriter, req *http.Request, fn PushHandler) error {
	branch, err := ParseWebhook(req, w.Secret)
	if err != nil {
		return err
	}
	go fn(branch)
	return nil
}
