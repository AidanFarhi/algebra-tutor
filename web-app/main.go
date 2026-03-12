package main

import (
	"io"
	"net/http"
)

func main() {

	mux := http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("http://ai-service:8000/api/v1/hello")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(body))
	})

	s := http.Server{
		Addr:    ":8090",
		Handler: &mux,
	}
	s.ListenAndServe()
}
