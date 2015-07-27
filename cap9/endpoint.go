package main

import (
	"log"
	"net/http"

	"go-in-action/cap9/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener: Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}