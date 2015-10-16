package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	webhook := NewWebhook()

	r := mux.NewRouter()
	r.Handle("/events", webhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	fmt.Println("Starting server on port", port)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
