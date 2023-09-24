package handler

import (
	"net/http"

	"github.com/aerostatas/interaction-go-kata/internal/api"
	"github.com/aerostatas/interaction-go-kata/internal/service"
)

type EventService interface {
	CreateEvent(service.EventCreate) (*service.Event, error)
}

type EventHandler struct {
	eventService EventService
}

func NewEventHandler(eventService EventService) *EventHandler {
	return &EventHandler{
		eventService,
	}
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event
// @Tags events
// @Accept json
// @Produce json
// @Param data body service.EventCreate true "Event data"
// @Success 201 {object} service.Event
// @Failure 400 {object} handler.ErrorResponse
// @Router /events [post]
func (h *EventHandler) CreateEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req service.EventCreate

		if err := api.ParseJSON(r, &req); err != nil {
			api.ErrorJSON(w, err)
			return
		}

		event, err := h.eventService.CreateEvent(req)
		if err != nil {
			api.ErrorJSON(w, err)
			return
		}

		api.ResponseJSON(w, event, http.StatusCreated)
	})
}
