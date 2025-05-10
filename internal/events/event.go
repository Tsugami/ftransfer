package events

import "time"

type Event struct {
	TransferId string
	ID         string
	Level      string
	Message    string
	CreatedAt  time.Time
}
