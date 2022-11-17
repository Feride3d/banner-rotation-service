package models

import (
	"time"
)

// Slot represents a slot.
type Slot struct {
	ID          string `db:"slot_id" json:"slotId"`
	Description string `db:"description" json:"description"`
}

// Banner represents a banner.
type Banner struct {
	ID          string `db:"banner_id" json:"bannerId"`
	Description string `db:"description" json:"description"`
}

// Group represents a socio-demographic group.
type Group struct {
	ID          string `db:"group_id" json:"groupId"`
	Description string `db:"description" json:"description"`
}

// Rotation represents which slot the banner is located in.
type Rotation struct {
	BannerID string `db:"banner_id" json:"bannerId"`
	SlotID   string `db:"slot_id" json:"slotId"`
}

// NewClick represents which group clicks on the banner in the slot.
type NewClick struct {
	BannerID string    `db:"banner_id" json:"bannerId"`
	SlotID   string    `db:"slot_id" json:"slotId"`
	GroupID  string    `db:"group_id" json:"groupId"`
	Time     time.Time `db:"time" json:"time"`
}

// BannerDisplay represents which slot and for which group the banner is displayed.
type BannerDisplay struct {
	BannerID string    `db:"banner_id" json:"bannerId"`
	SlotID   string    `db:"slot_id" json:"slotId"`
	GroupID  string    `db:"group_id" json:"groupId"`
	Time     time.Time `db:"time" json:"time"`
}

func (d *BannerDisplay) SetTime() {
	d.Time = time.Now().UTC()
}

// BannerNotDisplay represents which slot and for which group the banner is not displayed.
type BannerNotDisplay struct {
	BannerID string `db:"banner_id" json:"bannerId"`
	SlotID   string `db:"slot_id" json:"slotId"`
}
