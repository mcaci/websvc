package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type helloHandler struct{}

// func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!\n")
// }

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
	var jsonBody struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var name string = jsonBody.Name
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func reqLogger(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received a %q request from %q", r.Method, r.RequestURI)
		in.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/hello", reqLogger(&helloHandler{}))
	http.ListenAndServe(":8080", nil)
}
