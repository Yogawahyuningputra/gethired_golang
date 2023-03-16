package routes

import (
	"backend/handlers"
	"backend/packages/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func TodosRoutes(r *mux.Router) {
	TodosRepository := repositories.RepositoryTodos(mysql.DB)
	h := handlers.HandlerTodos(TodosRepository)

	r.HandleFunc("/todo-items", h.FindTodos).Methods("GET")
	r.HandleFunc("/todo-items/{id}", h.GetTodos).Methods("GET")
	r.HandleFunc("/todo-items", h.CreateTodos).Methods("POST")
	r.HandleFunc("/todo-items/{id}", h.UpdateTodos).Methods("PATCH")
	r.HandleFunc("/todo-items/{id}", h.DeleteTodos).Methods("DELETE")
}
