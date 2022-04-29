package models

import "github.com/Feride3d/banner-rotation-service/internal/service/mab"

// Slot represents a slot.
type Slot struct {
	ID          int
	Description string
}

// Banner represents a banner.
type Banner struct {
	ID          int
	Description string
}

// Group represents a socio-demographic group.
type Group struct {
	ID          int
	Description string
}

// Rotation represents a banner rotation.
type Rotation struct {
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
	BannerID int `db:"banner_id" json:"bannerId"`
	SlotID   int `db:"slot_id" json:"slotId"`
	GroupID  int `db:"group_id" json:"groupId"`
}

// DisplayBanner represents which slot and for which group the banner is displayed.
type DisplayBanner struct {
	BannerID int `db:"banner_id" json:"bannerId"`
	SlotID   int `db:"slot_id" json:"slotId"`
	GroupID  int `db:"group_id" json:"groupId"`
}

// NotDisplayBanner represents which slot and for which group the banner is not displayed.
type NotDisplayBanner struct {
	SlotID   string `db:"slot_id" json:"slotId"`
	BannerID string `db:"banner_id" json:"bannerId"`
}

type Storage interface {
	// Methods description
	Connect() error                                             // connect storage
	Close() error                                               // close storage
	AddBanner(bannerID, slotID int) error                       // add banner to slot
	DeleteBanner(bannerID, slotID int) error                    // delete banner from slot
	AddClick(bunnerID, slotID, groupID int) error               // add banner click
	AddBannerDisplay(slotID, groupID int) (mab.BannerID, error) // display banner
}
