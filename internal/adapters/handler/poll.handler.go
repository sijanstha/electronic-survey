package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

type pollHandler struct {
	pollService ports.PollService
}

func NewPollHandler(pollService ports.PollService) *pollHandler {
	return &pollHandler{pollService: pollService}
}

func (h *pollHandler) HandleSavePoll(w http.ResponseWriter, r *http.Request) error {
	createPollRequest := new(domain.CreatePollRequest)
	err := json.NewDecoder(r.Body).Decode(createPollRequest)
	if err != nil {
		return err
	}

	res, err := h.pollService.SavePoll(r.Context(), createPollRequest)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(res))
}
