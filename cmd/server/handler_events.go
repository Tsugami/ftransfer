package main

import (
	"net/http"
	"time"

	"github.com/Tsugami/ftransfer/internal/events"
	"github.com/gin-gonic/gin"
)

type EventResponse struct {
	ID         string    `json:"id"`
	TransferID string    `json:"transfer_id"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

func (e *EventResponse) FromEvent(event *events.Event) *EventResponse {
	return &EventResponse{
		ID:         event.ID,
		TransferID: event.TransferId,
		Level:      event.Level,
		Message:    event.Message,
		CreatedAt:  event.CreatedAt,
	}
}

func (s *Handler) ListEvents(c *gin.Context) {
	transferId := c.Query("transfer_id")
	if transferId == "" {
		c.JSON(http.StatusOK, []EventResponse{})
		return
	}
	events, err := s.eventService.ReadMany(c.Request.Context(), events.ReadManyOptions{
		TransferId: transferId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	response := []EventResponse{}
	for _, event := range events {
		eventResponse := &EventResponse{}
		response = append(response, *eventResponse.FromEvent(event))
	}

	c.JSON(http.StatusOK, response)
}
