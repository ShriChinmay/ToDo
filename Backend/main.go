package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var db=make(map[int] Task)
var nextId=1

func main(){
	R:=chi.NewRouter()
	R.Use(middleware.Logger)
	R.Use(middleware.Recoverer)
	R.Get("/todos", getTodos)
	R.Post("/todos", postTodo)
	R.Put("/todos/{id}", modifyTodo)
	R.Patch("/todos/{id}", modifyTodo)
	R.Delete("/todos/{id}", deleteTodo)
	R.Get("/todos/{id}", getTodo)
	err := http.ListenAndServe(":8080", R)
	if err != nil {
		panic(err)
	}
}