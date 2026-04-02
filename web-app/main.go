package main

import (
	"algtutor/controller"
	"algtutor/repo"
	"algtutor/service"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// create db pool
	connString := "postgres://postgres:postgres@db:5432/algtutor"
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("error creating db pool", err.Error())
	}

	aur := repo.NewAppUserRepo(pool)
	aus := service.NewAppUserService(aur)
	auc := controller.NewAppUserController(aus)

	mux := http.ServeMux{}

	mux.HandleFunc("/createuser", auc.CreateUser)

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("content-type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	resp, _ := http.Get("http://ai-service:8000/api/v1/hello")
	// 	defer resp.Body.Close()
	// 	body, _ := io.ReadAll(resp.Body)
	// 	w.Header().Add("content-type", "application/json")
	// 	w.Write([]byte(body))
	// })

	s := http.Server{
		Addr:    ":8090",
		Handler: &mux,
	}

	fmt.Println("server about to start")

	s.ListenAndServe()
}
