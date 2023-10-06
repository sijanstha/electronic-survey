package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type participantListHandler struct {
	participantListService ports.ParticipantListService
}

func NewParticipantListHandler(participantListService ports.ParticipantListService) *participantListHandler {
	return &participantListHandler{participantListService: participantListService}
}

func (h *participantListHandler) HandleParticipantList(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return h.handleSaveParticipantList(w, r)
	}

	return fmt.Errorf("%s method not supported", r.Method)
}

func (h *participantListHandler) handleSaveParticipantList(w http.ResponseWriter, r *http.Request) error {
	request := new(domain.CreateParticipantListRequest)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return err
	}

	res, err := h.participantListService.SaveParticipantList(r.Context(), request)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(res))
}
