package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Feride3d/banner-rotation-service/internal/storage"
	"github.com/Feride3d/banner-rotation-service/internal/storage/models"
)

type Service struct { //

}

// NewServer accesses storage.
func NewService(bd *storage.Storage) *Service {
	return &Service{}
}

// BannerRotationService structure.
type BannerRotationService struct {
	RotationStorage models.RotationStorage
	EventsStorage   models.EventsStorage
	EventsService   EventsService
}

// EventServ represents storage of events.
type EventServ struct {
	EventsStorage models.EventsStorage
}

// EventsService represents interface with method which saves events.
type EventsService interface {
	SaveEvents(ctx context.Context, rotation models.Rotation, groupID int, eventsType int) (*models.Events, error)
}

// AddBanner adds a banner to the slot.
func (b *BannerRotationService) AddBanner(ctx context.Context, rotation models.Rotation) (*models.Rotation, error) {
	newRotation, err := b.RotationStorage.AddBanner(ctx, rotation)
	if err != nil {
		return nil, fmt.Errorf("add banner to slot: %w", err)
	}

	return newRotation, nil
}

// DeleteBanner deletes the banner from the slot.
func (b *BannerRotationService) DeleteBanner(ctx context.Context, bannerID int) error {
	err := b.RotationStorage.DeleteBanner(ctx, bannerID)
	if err != nil {
		return fmt.Errorf("delete banner from slot: %w", err)
	}

	return nil
}

// CreateBanner creates a new banner.
func (b *BannerRotationService) CreateBanner(ctx context.Context, bannerID int, description string) error {
	err := b.RotationStorage.CreateBanner(ctx, bannerID, description)
	if err != nil {
		return fmt.Errorf("create banner: %w", err)
	}

	return nil
}

// CreateSlot creates a new slot.
func (b *BannerRotationService) CreateSlot(ctx context.Context, slotID int, description string) error {
	err := b.RotationStorage.CreateSlot(ctx, slotID, description)
	if err != nil {
		return fmt.Errorf("create slot: %w", err)
	}

	return nil
}

// CreateGroup creates a new group.
func (b *BannerRotationService) CreateGroup(ctx context.Context, groupID int, description string) error {
	err := b.RotationStorage.CreateGroup(ctx, groupID, description)
	if err != nil {
		return fmt.Errorf("create group: %w", err)
	}

	return nil
}

// AddClick adds a click on the banner in the group.
func (b *BannerRotationService) AddClick(ctx context.Context, rotation models.Rotation, groupID int,
) (*models.Events, error) {
	events, err := b.EventsService.SaveEvents(ctx, rotation, groupID, models.EventsClickType)
	if err != nil {
		return nil, fmt.Errorf("add a click on banner: %w", err)
	}

	return events, nil
}

// Adds a banner to display.
func (b *BannerRotationService) AddBannerDisplay(ctx context.Context, slotID int, groupID int,
) (int, *models.Events, error) {
	rotations, err := b.RotationStorage.GetAllBannersInSlot(ctx, slotID)
	if err != nil {
		return 0, nil, fmt.Errorf("get banners by slotID: %w", err)
	}

	eventsList, err := b.EventsStorage.FindAll(ctx, slotID, groupID)
	if err != nil {
		return 0, nil, fmt.Errorf("find events by slotID and groupID: %w", err)
	}

	rotation, err := b.chooseBanner(rotations, eventsList)
	if err != nil {
		return 0, nil, fmt.Errorf("define banner: %w", err)
	}

	events, err := b.EventsService.SaveEvents(ctx, *rotation, groupID, models.EventsDisplayType)
	if err != nil {
		return 0, nil, fmt.Errorf("save display event: %w", err)
	}

	return rotation.BannerID, events, nil
}

// ChooseBanner chooses a banner to display.
func (b *BannerRotationService) chooseBanner(rotations []*models.Rotation, eventsList []*models.Events,
) (*models.Rotation, error) {
	if len(rotations) <= 0 {
		return nil, errors.New("empty list")
	}

	banners := make(map[int]models.Banner, len(rotations))
	for _, rotation := range rotations {
		banners[rotation.BannerID] = models.Banner{ID: rotation.BannerID}
	}

	for _, events := range eventsList {
		banner, check := banners[events.BannerID]

		if !check {
			continue
		}

		if events.DisplayType() {
			banner.Displays++
		}

		if events.ClickType() {
			banner.Clicks++
		}

		banner.GroupID = events.GroupID

		banners[banner.ID] = banner
	}

	display := make([]int, 0, len(banners))
	click := make([]int, 0, len(banners))
	mab := make(map[int]int, len(banners))
	i := 0

	for bannerID, banner := range banners {
		mab[i] = bannerID
		display = append(display, banner.Displays)
		click = append(click, banner.Clicks)
		i++
	}
	rotation := new(models.Rotation)

	for _, r := range rotations {
		{
			rotation = r
		}
	}

	return rotation, nil
}

// SaveEvents save events for statistical information.
func (e *EventServ) SaveEvents(ctx context.Context, rotation models.Rotation, groupID int, eventsType int,
) (*models.Events, error) {
	events := models.Events{
		Type:     eventsType,
		BannerID: rotation.BannerID,
		SlotID:   rotation.SlotID,
		GroupID:  groupID,
		Date:     time.Now().UTC(),
	}

	newEvents, err := e.EventsStorage.Add(ctx, events)
	if err != nil {
		return nil, fmt.Errorf("save events: %w", err)
	}

	return newEvents, nil
}
