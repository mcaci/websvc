package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mcaci/draft-websvc/proverbs"
)

func incomingReqLoggerMdw(in http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("received a %q request from %q", r.Method, r.RequestURI)
			if r == nil || r.Body == nil {
				http.Error(w, "Request is empty", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			log.Printf("the body of the request is: %q", string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body))
			in.ServeHTTP(w, r)
		})
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil || r.Body == nil {
		http.Error(w, "Request is empty", http.StatusBadRequest)
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
	name := "World"
	if v.Name != "" {
		name = v.Name
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", incomingReqLoggerMdw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello!") })))
	mux.Handle("/greet", incomingReqLoggerMdw(&helloHandler{}))
	mux.Handle("/proverbs/", incomingReqLoggerMdw(&proverbs.Handler{}))

	err := http.ListenAndServe(":8080", mux)
	log.Println(err)
}
