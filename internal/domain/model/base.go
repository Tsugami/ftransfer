package model

import (
	"time"
)

// Base contains common fields for all domain models
type Base struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tags      Tags
}

// NewBase creates a new base with current timestamps
func NewBase() Base {
	now := time.Now()
	return Base{
		CreatedAt: now,
		UpdatedAt: now,
		Tags:      make(Tags, 0),
	}
}

// UpdateTimestamp updates the UpdatedAt field to current time
func (b *Base) UpdateTimestamp() {
	b.UpdatedAt = time.Now()
}

// AddTag adds a new tag
func (b *Base) AddTag(name, value string) {
	AddTag(&b.Tags, name, value)
}

// GetTagValue returns the value of a tag by name
func (b *Base) GetTagValue(name string) (string, bool) {
	if tag, found := FindByName(b.Tags, name); found {
		return tag.Value, true
	}
	return "", false
}

// RemoveTag removes a tag by name
func (b *Base) RemoveTag(name string) {
	RemoveTag(&b.Tags, name)
}
