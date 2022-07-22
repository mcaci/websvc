package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
- func main() and write http.ListenAndServe
- log error returned
- explain port and http handler
- create a type (hellohandler)
- implement ServeHTTP method // n, err := w.Write([]byte(msg)) || fmt.Fprintln(w, msg)
- RUN THE EXAMPLE
- add simple hello function with same body
- RUN THE EXAMPLE
- explain the http.HandlerFunc
*/

/* CONCEPTS
- multi error return
- implementing an interface
*/

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	n, err := fmt.Fprintln(w, msg)
	log.Println(n, err)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	n, err := fmt.Fprintln(w, msg)
	log.Println(n, err)
}

func main() {
	// var h http.Handler = &helloHandler{}
	var h http.Handler = http.HandlerFunc(helloWorldHandler)
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}
