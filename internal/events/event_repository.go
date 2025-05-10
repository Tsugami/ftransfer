package events

import "context"

type EventRepository interface {
	CreateMany(ctx context.Context, events []*Event) error
	ReadMany(ctx context.Context, transferId string) ([]*Event, error)
}
