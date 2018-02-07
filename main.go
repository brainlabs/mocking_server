package main

import (
	"log"
	"github.com/brainlabs/mocking_server/route"
	"net/http"
)

func main() {
	r := new(route.Router).Make()
	log.Fatal(http.ListenAndServe(":9191", r))
}
