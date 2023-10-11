package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	if r.Method == "PUT" {
		return h.handleUpdateParticipantList(w, r)
	}

	if r.Method == "GET" {
		return h.handleGetParticipantList(w, r)
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

func (h *participantListHandler) handleUpdateParticipantList(w http.ResponseWriter, r *http.Request) error {
	request := new(domain.UpdateParticipantList)
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return err
	}

	res, err := h.participantListService.UpdateParticipantList(r.Context(), request)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(res))
}

func (h *participantListHandler) handleGetParticipantList(w http.ResponseWriter, r *http.Request) error {
	paginationFilter := domain.ParsePaginationRequest(r)
	log.Println(paginationFilter.Page, paginationFilter.Limit, paginationFilter.Sort, paginationFilter.SortBy)

	filter := domain.ParticipantListFilter{
		PaginationFilter: paginationFilter,
	}

	resp, err := h.participantListService.GetParticipantList(r.Context(), &filter)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}

func (h *participantListHandler) HandleGetParticipantListById(w http.ResponseWriter, r *http.Request) error {
	strID := mux.Vars(r)["id"]
	id := utils.ParseInteger(strID)

	pl, err := h.participantListService.GetParticipantListById(r.Context(), int64(id))
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(pl))
}
