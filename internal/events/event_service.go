package events

import "context"

type EventService struct {
	repository EventRepository
}

func NewEventService(repository EventRepository) *EventService {
	return &EventService{repository: repository}
}

type ReadManyOptions struct {
	TransferId string
}

func (s *EventService) ReadMany(ctx context.Context, options ReadManyOptions) ([]*Event, error) {
	return s.repository.ReadMany(ctx, options.TransferId)
}
