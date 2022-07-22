package main

import (
	"log"
	"net/http"

	"github.com/mcaci/websvc/drafts/03-routes/hello"
	"github.com/mcaci/websvc/proverbs"
)

/*
- func main() and write with Handler http.ListenAndServe
- log error returned
- create http.Handler with hello.Handler
- create another http.Handler with proverbs.Handler
- change second parameter from nil to any ServeMux name
- EXPLAIN the usage of nil values as defaults
- add the two handle muxes: (remember the / at the end of proverbs)
	mux.Handle("/hello", h)
	mux.Handle("/proverbs/", p)
*/

/* CONCEPTS
- multiplexer definition
- make the zero value useful
*/

func main() {
	var h http.Handler = &hello.Handler{}
	var p http.Handler = &proverbs.Handler{}

	mux := http.NewServeMux()
	mux.Handle("/hello", h)
	mux.Handle("/proverbs/", p)

	err := http.ListenAndServe(":8080", mux)
	log.Println(err)
}
