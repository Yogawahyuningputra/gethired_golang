package handlers

import (
	activitydto "backend/dto/activity"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerActivity struct {
	ActivityRepository repositories.ActivityRepository
}

func HandlerActivity(ActivityRepository repositories.ActivityRepository) *handlerActivity {
	return &handlerActivity{ActivityRepository}
}

func (h *handlerActivity) FindActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	activities, err := h.ActivityRepository.FindActivity()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activities}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerActivity) GetActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	activity, err := h.ActivityRepository.GetActivity(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Activity with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: activity}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerActivity) CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(activitydto.ActivityRequest)
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

	activity := models.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	data, err := h.ActivityRepository.CreateActivity(activity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Bad Request", Message: "title cannot be null"}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerActivity) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(activitydto.ActivityUpdate)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Request Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	activity, err := h.ActivityRepository.GetActivity(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Activity with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		activity.Title = request.Title
	}

	data, err := h.ActivityRepository.UpdateActivity(activity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Bad Request", Message: "Failed"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerActivity) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	activity, err := h.ActivityRepository.GetActivity(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Activity with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ActivityRepository.DeleteActivity(activity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Not Found", Message: fmt.Sprintf("Activity with ID %s not found", strconv.Itoa(id))}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Message: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}
