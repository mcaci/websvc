package proverbs

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	proverbs []proverb
}

func (ph *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[len("/proverbs/"):])
	if err != nil {
		fmt.Fprintln(w, "no proverb provided: choosing at random")
	}
	ps, err := load("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if id >= len(ps) {
		http.Error(w, "Proverb not found", http.StatusNotFound)
		return
	}
	switch id {
	case 0:
		rand.Seed(time.Now().UnixMilli())
		id = rand.Intn(len(ps))
		fmt.Fprint(w, ps[id].String())
	default:
		fmt.Fprint(w, ps[id-1].String())
	}
}
