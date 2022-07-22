package hello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
