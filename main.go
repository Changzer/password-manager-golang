package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"password-manager/configs"
	"password-manager/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	r.Get("/done", handlers.GetByDone)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
