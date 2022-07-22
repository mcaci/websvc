package mdw

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

/*
- explain a middleware in http
- write a func with http.Handler as param and return type
- return an http.HandlerFunc casting a func(w http.ResponseWriter, r *http.Request)
- copy the initial part of the previous example
- add the two log lines
-- log.Printf("Incoming request from %q", r.RequestURI)
-- log.Printf("Request parameters: %q", string(body))
- add the clone of the req with io.NopCloser(bytes.NewBuffer(body)) body -> []byte after ReadAll
- call the ServeHTTP
*/

/* CONCEPTS
- returning a function (first class elements)
- mdw: http.Handler wrapping another http.Handler
*/

func ReqMonitor(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming %s request from %q", r.Method, r.RequestURI)
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
		log.Printf("Request parameters: %q", string(body))

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		in.ServeHTTP(w, r)
	})
}
