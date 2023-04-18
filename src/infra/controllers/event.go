package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/matheusteodoro01/go-clean-arch/src/domain/usecase"
)

type SendEventController struct {
	SendEventUseCase *usecase.SendEventUseCase
}

func NewEventController(sendEventUseCase *usecase.SendEventUseCase) *SendEventController {
	return &SendEventController{SendEventUseCase: sendEventUseCase}
}

func (controller *SendEventController) SendEventController(w http.ResponseWriter, r *http.Request) {

	var input usecase.SendEventInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.SendEventUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Event sent")
}
