package routes

import (
	"backend/handlers"
	"backend/packages/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func ActivityRoutes(r *mux.Router) {
	ActivityRepository := repositories.RepositoryActivity(mysql.DB)
	h := handlers.HandlerActivity(ActivityRepository)

	r.HandleFunc("/activity-groups", h.FindActivity).Methods("GET")
	r.HandleFunc("/activity-groups/{id}", h.GetActivity).Methods("GET")
	r.HandleFunc("/activity-groups", h.CreateActivity).Methods("POST")
	r.HandleFunc("/activity-groups/{id}", h.UpdateActivity).Methods("PATCH")
	r.HandleFunc("/activity-groups/{id}", h.DeleteActivity).Methods("DELETE")
}
