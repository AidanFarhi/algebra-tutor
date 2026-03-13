package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestDBConnection(connString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return fmt.Errorf("failed to create pool: %w", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	fmt.Println("✅ Successfully connected to Postgres")
	return nil
}

func main() {

	// attempt to connect to db
	connString := "postgres://postgres:postgres@db:5432/algtutor"
	TestDBConnection(connString)

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

	fmt.Println("server about to start")

	s.ListenAndServe()
}
