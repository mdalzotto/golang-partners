package handlers

import (
	"desafio/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	usecase *usecase.EventUseCases
}

type ReserveRequest struct {
	Spots []string `json:"spots"`
}

func NewEventHandler(uc *usecase.EventUseCases) *EventHandler {
	return &EventHandler{usecase: uc}
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events := h.usecase.GetEvents()
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["eventId"])

	event, err := h.usecase.GetEventByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) GetEventSpots(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["eventId"])

	spots, err := h.usecase.GetSpotsByEventID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(spots)
}

func (h *EventHandler) ReserveSpot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventID, _ := strconv.Atoi(params["eventId"])

	var req ReserveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	spot, err := h.usecase.ReserveSpot(eventID, req.Spots)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(spot)
}
