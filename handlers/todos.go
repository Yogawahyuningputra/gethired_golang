package handlers

import (
	dto "backend/dto/result"
	todosdto "backend/dto/todos"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTodos struct {
	TodosRepository repositories.TodosRepository
}

func HandlerTodos(TodosRepository repositories.TodosRepository) *handlerTodos {
	return &handlerTodos{TodosRepository}
}

func (h *handlerTodos) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := h.TodosRepository.FindTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: convertResponse(todos)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTodos) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	todo, err := h.TodosRepository.GetTodos(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Todos with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: convertResponse(todo)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTodos) CreateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(todosdto.TodosRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Bad Request", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "field cannot be null", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	todo := models.Todos{
		ActivityGroupID: request.ActivityGroupID,
		Title:           request.Title,
		IsActive:        request.IsActive,
		Priority:        "very-high",
	}

	data, err := h.TodosRepository.CreateTodos(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Bad Request", Message: "title cannot be null"}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTodos) UpdateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(todosdto.TodosUpdate)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Request Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := h.TodosRepository.GetTodos(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Todo with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Priority != "" {
		todo.Priority = request.Priority
	}

	if request.IsActive != todo.IsActive {
		todo.IsActive = request.IsActive
	}

	data, err := h.TodosRepository.UpdateTodos(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Bad Request", Message: "Failed"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTodos) DeleteTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := h.TodosRepository.GetTodos(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Todo with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = h.TodosRepository.DeleteTodos(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Todo with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: make(map[string]interface{})}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u interface{}) interface{} {
	switch u := u.(type) {
	case models.Todos:
		return todosdto.TodosResponse{
			TodoID:          u.TodoID,
			ActivityGroupID: u.ActivityGroupID,
			Title:           u.Title,
			IsActive:        u.IsActive,
			Priority:        u.Priority,
			CreatedAt:       u.CreatedAt,
			UpdatedAt:       u.UpdatedAt,
		}
	case []models.Todos:
		todosResponse := make([]todosdto.TodosResponse, len(u))
		for i, todo := range u {
			todosResponse[i] = todosdto.TodosResponse{
				TodoID:          todo.TodoID,
				ActivityGroupID: todo.ActivityGroupID,
				Title:           todo.Title,
				IsActive:        todo.IsActive,
				Priority:        todo.Priority,
				CreatedAt:       todo.CreatedAt,
				UpdatedAt:       todo.UpdatedAt,
			}
		}
		return todosResponse
	default:
		return nil
	}
}
