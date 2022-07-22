package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mcaci/websvc/proverbs"
)

func mdw(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received a %q request from the service at %q", r.Method, r.RequestURI)
		if r == nil || r.Body == nil {
			http.Error(w, "request is not valid", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading the body", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPost {
			log.Printf("the body of the request was %q", string(body))
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		in.ServeHTTP(w, r)
	})
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil || r.Body == nil {
		http.Error(w, "request is not valid", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the body", http.StatusBadRequest)
		return
	}
	var jsonReq struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	err1 := json.Unmarshal(body, &jsonReq)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	name := jsonReq.Name
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main() {
	mux := http.NewServeMux()
	h := http.Handler(&helloHandler{})
	mux.Handle("/greet", h)
	mux.Handle("/proverbs/", &proverbs.Handler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Sorry I can't handle this case\n") })

	err := http.ListenAndServe(":8080", mdw(mux))
	log.Println(err)
}
