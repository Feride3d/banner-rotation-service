package models

import (
	"context"
	"time"
)

// Slot represents a slot.
type Slot struct {
	ID          int
	Description string
}

// Banner represents a banner.
type Banner struct {
	ID          int
	Description string
	Displays    int
	Clicks      int
	GroupID     int
}

// Group represents a socio-demographic group.
type Group struct {
	ID          int
	Description string
}

// Rotation represents a banner rotation.
type Rotation struct {
	ID       int
	BannerID int
	SlotID   int
	GroupID  int
	Display  int
	Click    int
}

// BannerLocation represents which slot the banner is located in.
type BannerLocation struct {
	BannerID int `db:"banner_id" json:"bannerId"`
	SlotID   int `db:"slot_id" json:"slotId"`
}

// NewClick represents which group clicks on the banner in the slot.
type NewClick struct {
	BannerID int       `db:"banner_id" json:"bannerId"`
	SlotID   int       `db:"slot_id" json:"slotId"`
	GroupID  int       `db:"group_id" json:"groupId"`
	Time     time.Time `db:"time" json:"time"`
}

// DisplayBanner represents which slot and for which group the banner is displayed.
type DisplayBanner struct {
	BannerID int       `db:"banner_id" json:"bannerId"`
	SlotID   int       `db:"slot_id" json:"slotId"`
	GroupID  int       `db:"group_id" json:"groupId"`
	Time     time.Time `db:"time" json:"time"`
}

func (d *DisplayBanner) SetTime() {
	d.Time = time.Now().UTC()
}

// NotDisplayBanner represents which slot and for which group the banner is not displayed.
type NotDisplayBanner struct {
	SlotID   string `db:"slot_id" json:"slotId"`
	BannerID string `db:"banner_id" json:"bannerId"`
}

// RotationStorage represents baneer rotation storage interface.
type RotationStorage interface {
	Connect() error
	Close() error
	CreateBanner(ctx context.Context, bannerID int, description string) error
	CreateSlot(ctx context.Context, slotID int, description string) error
	CreateGroup(ctx context.Context, groupID int, description string) error
	AddBanner(ctx context.Context, rotation Rotation) (*Rotation, error)
	GetBanner(ctx context.Context, bannerID int) (*Rotation, error)
	GetAllBannersInSlot(ctx context.Context, slotID int) ([]*Rotation, error)
	DeleteBanner(ctx context.Context, bannerID int) error
}
