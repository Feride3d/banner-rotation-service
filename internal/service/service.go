package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Feride3d/banner-rotation-service/internal/storage"
	"github.com/Feride3d/banner-rotation-service/internal/storage/models"
)

type Service struct {
	logger  Logger
	storage storage.RotationStorage
	mab     Mab
	queues  Queues
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
	Trace(msg string)
}

type Mab interface {
	Ucb(banners []string, clicks map[string]int, displays map[string]int) (string, error)
}

type Queues interface {
	Publish(ctx context.Context, event models.Events) error
}

// Logger
func (s *Service) LoggerBegin() Logger {
	return s.logger
}

// NewService
func New(logger Logger, storage storage.RotationStorage, mab Mab, queues Queues) *Service {
	return &Service{logger: logger, storage: storage, mab: mab, queues: queues}
}

// AddBanner adds a banner to the slot.
func (s *Service) AddBanner(ctx context.Context, bannerID string, slotID string) error {
	err := s.storage.AddBanner(ctx, bannerID, slotID)
	if err != nil {
		return fmt.Errorf("add banner to slot: %w", err)
	}

	return nil
}

// DeleteBanner deletes the banner from the slot.
func (s *Service) DeleteBanner(ctx context.Context, bannerID string, slotID string) error {
	err := s.storage.DeleteBanner(ctx, bannerID, slotID)
	if err != nil {
		return fmt.Errorf("delete banner from slot: %w", err)
	}

	return nil
}

// CreateBanner creates a new banner.
func (s *Service) CreateBanner(ctx context.Context, bannerID string, description string) (string, error) {
	_, err := s.storage.CreateBanner(ctx, bannerID, description)
	if err != nil {
		return "", fmt.Errorf("create banner: %w", err)
	}

	return "", nil
}

// CreateSlot creates a new slot.
func (s *Service) CreateSlot(ctx context.Context, slotID string, description string) (string, error) {
	_, err := s.storage.CreateSlot(ctx, slotID, description)
	if err != nil {
		return "", fmt.Errorf("create slot: %w", err)
	}

	return "", nil
}

// CreateGroup creates a new group.
func (s *Service) CreateGroup(ctx context.Context, groupID string, description string) (string, error) {
	_, err := s.storage.CreateGroup(ctx, groupID, description)
	if err != nil {
		return "", fmt.Errorf("create group: %w", err)
	}

	return "", nil
}

// AddClick adds a click on the banner in the group.
func (s *Service) AddClick(ctx context.Context, bannerID string, slotID string, groupID string) error {
	time := time.Now()
	err := s.storage.AddClick(ctx, bannerID, slotID, groupID, time)
	if err != nil {
		return fmt.Errorf("add a click on banner: %w", err)
	}

	err = s.queues.Publish(ctx, models.Events{Type: "click", BannerID: bannerID, SlotID: slotID, GroupID: groupID, Time: time})
	if err != nil {
		return fmt.Errorf("failed to create click event, %w", err)
	}
	return nil
}

// Adds a banner to display.
func (s *Service) AddDisplay(ctx context.Context, bannerID, slotID string, groupID string) error {
	time := time.Now()
	err := s.storage.AddBannerDisplay(ctx, bannerID, slotID, groupID, time)
	if err != nil {
		return fmt.Errorf("add banner's display: %w", err)
	}

	err = s.queues.Publish(ctx, models.Events{Type: "display", BannerID: bannerID, SlotID: slotID, GroupID: groupID, Time: time})
	if err != nil {
		return fmt.Errorf("publish banner display: %w", err)
	}

	return nil
}

// AddBannerDisplay chooses a banner to display.
func (s *Service) AddBannerDisplay(ctx context.Context, slotID string, groupID string) (string, error) {
	bannerNotDisplay, err := s.storage.GetBannerNotDisplay(ctx, slotID)
	if err != nil {
		return "", fmt.Errorf("choose banner for display: %w", err)
	}

	if len(bannerNotDisplay) > 0 {
		bannerID := bannerNotDisplay[0].BannerID

		err := s.AddDisplay(ctx, bannerID, slotID, groupID)
		if err != nil {
			return "", fmt.Errorf("choose banner for display: %w", err)
		}
		return bannerID, nil
	}

	bannerInSlot, err := s.storage.GetBannerInSlot(ctx, slotID)
	if err != nil {
		return "", err
	}

	bannerClick, err := s.storage.GetBannerClick(ctx, slotID)
	if err != nil {
		return "", err
	}

	bannerDisplay, err := s.storage.GetBannerDisplay(ctx, slotID)
	if err != nil {
		return "", err
	}

	banner, bannerClickInMap, bannerDisplayInMap := s.CountInMap(bannerInSlot, bannerClick, bannerDisplay)
	bannerID, err := s.mab.Ucb(banner, bannerClickInMap, bannerDisplayInMap)
	if err != nil {
		return "", err
	}

	err = s.AddDisplay(ctx, bannerID, slotID, groupID)
	if err != nil {
		return "", err
	}

	return bannerID, nil
}

func (s *Service) CountInMap(
	bannerInSlot []models.Rotation,
	bannerClick []models.NewClick,
	bannerDisplay []models.BannerDisplay) (
	banner []string,
	bannerClickInMap map[string]int,
	bannerDisplayInMap map[string]int,
) {
	bannerClickInMap = make(map[string]int)
	bannerDisplayInMap = make(map[string]int)
	for _, banners := range bannerInSlot {
		banner = append(banner, banners.BannerID)
	}
	for _, click := range bannerClick {
		bannerClickInMap[click.BannerID]++
	}
	for _, display := range bannerDisplay {
		bannerDisplayInMap[display.BannerID]++
	}

	return banner, bannerClickInMap, bannerDisplayInMap

}
