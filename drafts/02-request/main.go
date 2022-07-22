package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
- func main() and write with Handler http.ListenAndServe
- log error returned
- create http.HandlerFunc function helloHandler
- check if request and request body exists + err return StatusBadRequest (don't forget the return statement)
- defer r.Body.Close
- call io.ReadAll on r.Body and check err
- call json.Unmarshall on anon struct{ Name string `json:"name"`}
- switch on Name value and print to responsewriter
*/

/* CONCEPTS
- if/else and switch vs readability
- struct internals
*/

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil || r.Body == nil {
		http.Error(w, "Empty request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var v struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch v.Name {
	case "":
		fmt.Fprintln(w, "Hello World!")
	default:
		fmt.Fprintf(w, "Hello %s!\n", v.Name)
	}
}

func main() {
	var h http.Handler = &helloHandler{}
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}
