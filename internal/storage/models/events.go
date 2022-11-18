package models

import (
	"context"
	"time"
)

var EventsDisplayType = 1

var EventsClickType = 2

type Events struct {
	Type     string    `db:"type" json:"type"`
	BannerID string    `db:"banner_id" json:"bannerId"`
	SlotID   string    `db:"slot_id" json:"slotId"`
	GroupID  string    `db:"group_id" json:"groupId"`
	Time     time.Time `db:"time" json:"time"`
}

// EventsStorage represents statistical information.
type EventsStorage interface {
	Add(ctx context.Context, events Events) (*Events, error)                 // add events
	FindAll(ctx context.Context, slotID int, groupID int) ([]*Events, error) // find all events in slot/group
	Delete(ctx context.Context, ID int) error                                // delete events
}
