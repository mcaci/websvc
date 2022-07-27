package main

// func reqLogger(in http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("received a %q request from %q", r.Method, r.RequestURI)
// 		in.ServeHTTP(w, r)
// 	})
// }

// type helloHandler struct{}
// func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!\n") }

func main() {
	// http.Handle("/greet", &helloHandler{})
	// http.ListenAndServe(":8080", nil)
}
