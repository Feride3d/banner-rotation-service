package models

import (
	"context"
	"time"
)

var EventsDisplayType = 1

var EventsClickType = 2

type Events struct {
	Type     int       `db:"type" json:"type"`
	BannerID int       `db:"banner_id" json:"banner_Id"`
	SlotID   int       `db:"slot_id" json:"slot_Id"`
	GroupID  int       `db:"group_id" json:"group_Id"`
	Date     time.Time `db:"date" json:"date"`
}

// EventsStorage represents statistical information.
type EventsStorage interface {
	Add(ctx context.Context, events Events) (*Events, error)                 // add events
	FindAll(ctx context.Context, slotID int, groupID int) ([]*Events, error) // find all events in slot/group
	Delete(ctx context.Context, ID int) error                                // delete events
}

func (e *Events) DisplayType() bool {
	return e.Type == EventsDisplayType
}

func (e *Events) ClickType() bool {
	return e.Type == EventsClickType
}
