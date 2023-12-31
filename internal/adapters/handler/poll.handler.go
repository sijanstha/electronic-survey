package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

func (h *pollHandler) HandlePoll(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return h.handleSavePoll(w, r)
	}

	if r.Method == "PUT" {
		return h.handleUpdatePoll(w, r)
	}

	if r.Method == "GET" {
		return h.handleGetPoll(w, r)
	}

	return fmt.Errorf("%s method not supported", r.Method)
}

func (h *pollHandler) handleSavePoll(w http.ResponseWriter, r *http.Request) error {
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

func (h *pollHandler) handleUpdatePoll(w http.ResponseWriter, r *http.Request) error {
	updatePollRequest := new(domain.UpdatePollRequest)
	err := json.NewDecoder(r.Body).Decode(updatePollRequest)
	if err != nil {
		return err
	}

	res, err := h.pollService.UpdatePoll(r.Context(), updatePollRequest)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(res))
}

func (h *pollHandler) handleGetPoll(w http.ResponseWriter, r *http.Request) error {
	paginationFilter := domain.ParsePaginationRequest(r)
	state := r.URL.Query().Get("state")
	showOwnPoll := r.URL.Query().Get("showOwnPoll")

	log.Println(paginationFilter.Page, paginationFilter.Limit, paginationFilter.Sort,
		paginationFilter.SortBy, state, showOwnPoll)

	states := make([]domain.PollState, 0)
	if len(state) > 0 {
		match, _ := regexp.MatchString("\\(([^)]+)\\)", state)
		if match {
			state = strings.TrimPrefix(state, "(")
			state = strings.TrimSuffix(state, ")")
			if strings.Contains(state, ",") {
				stateSlice := strings.Split(state, ",")
				for _, data := range stateSlice {
					states = append(states, domain.PollState(data))
				}
			} else {
				states = append(states, domain.PollState(state))
			}
			log.Println(states)
		}
	}

	filter := domain.PollListFilter{
		PaginationFilter:   paginationFilter,
		FilterPrimaryOwner: utils.ParseBoolean(showOwnPoll),
		States:             states,
	}

	resp, err := h.pollService.GetAllPoll(r.Context(), filter)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}

func (h *pollHandler) HandleGetPollById(w http.ResponseWriter, r *http.Request) error {
	strID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	poll, err := h.pollService.GetPollById(int64(id))
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, domain.NewApiResponse(poll))
}
