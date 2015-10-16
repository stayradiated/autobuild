package main

import (
	"fmt"
	"net/http"
)

type Webhook struct {
}

func NewWebhook() *Webhook {
	return &Webhook{}
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
	fmt.Println(branch)
}
