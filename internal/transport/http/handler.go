package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/makaksel/MRNotifier/internal/domain"
	"github.com/makaksel/MRNotifier/internal/usecase"
)

type Handler struct {
	useCase *usecase.Client
}

func NewHandler(uc *usecase.Client) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) CreateMR(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateMRRequest

	log.Printf("HANDLER CALLED")

	json.NewDecoder(r.Body).Decode(&req)

	err := h.useCase.HandleMr(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
