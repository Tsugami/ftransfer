package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Tsugami/ftransfer/internal/events"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateMany(ctx context.Context, events []*events.Event) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO events (transfer_id, level, message, created_at)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, event := range events {
		_, err = stmt.ExecContext(ctx, event.TransferId, event.Level, event.Message, event.CreatedAt)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *EventRepository) ReadMany(ctx context.Context, transferId string) ([]*events.Event, error) {
	fmt.Println("ReadMany", transferId)
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, transfer_id, level, message, created_at
		FROM events
		WHERE transfer_id = $1
		ORDER BY created_at DESC
		LIMIT 30
	`, transferId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*events.Event
	for rows.Next() {
		event := &events.Event{}
		err = rows.Scan(&event.ID, &event.TransferId, &event.Level, &event.Message, &event.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, event)
	}

	return result, nil
}
