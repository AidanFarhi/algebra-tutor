package main

import "net/http"

func main() {
	mux := http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	s := http.Server{
		Addr:    ":8090",
		Handler: &mux,
	}
	s.ListenAndServe()
}
