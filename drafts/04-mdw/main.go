package main

import (
	"log"
	"net/http"

	"github.com/mcaci/drafts/04-mdw/hello"
	"github.com/mcaci/drafts/04-mdw/mdw"
	"github.com/mcaci/drafts/proverbs"
)

/*
- func main() and write with Handler http.ListenAndServe
- log error returned
- create http.HandlerFunc function helloHandler
- check if request and request body exists + err return StatusBadRequest (don't forget the return statement)
- defer r.Body.Close
- call io.ReadAll on r.Body and check err
- call json.Unmarshall on anon struct{ Name string `json:"name"`}
- revert and use the request
*/

/* CONCEPTS
- if/else and switch vs readability
- struct internals
*/

func main() {
	h := mdw.ReqMonitor(&hello.Handler{})
	p := mdw.ReqMonitor(&proverbs.Handler{})

	mux := http.NewServeMux()
	mux.Handle("/proverbs/", p)
	mux.Handle("/hello", h)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("noop\n")) })

	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}
