package events

import (
	"context"
	"fmt"
	"time"
)

type EventLogger struct {
	repository EventRepository
	transferId string
	data       []*Event
}

func NewEventLogger(transferId string, repository EventRepository) *EventLogger {
	return &EventLogger{transferId: transferId, repository: repository}
}

func (l *EventLogger) Info(ctx context.Context, message string) {
	fmt.Println("INFO", message)
	l.data = append(l.data, &Event{
		Level:      "info",
		Message:    message,
		TransferId: l.transferId,
		CreatedAt:  time.Now(),
	})
}

func (l *EventLogger) Error(ctx context.Context, message string) {
	fmt.Println("ERROR", message)
	l.data = append(l.data, &Event{
		Level:      "error",
		Message:    message,
		TransferId: l.transferId,
		CreatedAt:  time.Now(),
	})
}

func (l *EventLogger) Debug(ctx context.Context, message string) {
	fmt.Println("DEBUG", message)
	l.data = append(l.data, &Event{
		Level:      "debug",
		Message:    message,
		TransferId: l.transferId,
		CreatedAt:  time.Now(),
	})
}

func (l *EventLogger) Flush(ctx context.Context) error {
	return l.repository.CreateMany(ctx, l.data)
}
