package main

import (
	"net/http"
)

type PushHandler func(branch string)

type Webhook struct {
	HandlePush PushHandler
}

func (w *Webhook) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		w.handlePOST(resp, req)
	}
}

func (w *Webhook) handlePOST(resp http.ResponseWriter, req *http.Request) {
	branch, err := ParseWebhook(req.Body)
	if err != nil {
		panic(err)
	}
	go w.HandlePush(branch)
}
