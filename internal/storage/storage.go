package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Feride3d/banner-rotation-service/internal/logger"
	"github.com/Feride3d/banner-rotation-service/internal/storage/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// RotationStorage represents baneer rotation storage interface.
type RotationStorage interface {
	CreateBanner(ctx context.Context, bannerID string, description string) (string, error)
	CreateSlot(ctx context.Context, slotID string, description string) (string, error)
	CreateGroup(ctx context.Context, groupID string, description string) (string, error)
	AddBanner(ctx context.Context, bannerID string, slotID string) error
	DeleteBanner(ctx context.Context, bannerID string, slotID string) error
	AddClick(ctx context.Context, bannerID string, slotID string, groupID string, time time.Time) error
	AddBannerDisplay(ctx context.Context, bannerID, slotID string, groupID string, time time.Time) error
	GetBannerDisplay(ctx context.Context, slotID string) (bannerDisplay []models.BannerDisplay, err error)
	GetBannerNotDisplay(ctx context.Context, slotID string) (bannerNotDisplay []models.BannerNotDisplay, err error)
	GetBannerClick(ctx context.Context, slotID string) (newClicks []models.NewClick, err error)
	GetBannerInSlot(ctx context.Context, slotID string) (rotation []models.Rotation, err error)
}

type Storage struct {
	conn   *pgxpool.Pool
	logger logger.Logger
}

// Connect connects to database
func (s *Storage) Connect(ctx context.Context, connStr string) (*Storage, error) {
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &Storage{conn: conn}, nil
}

// Close closes connection to database.
func (s *Storage) Close(ctx context.Context) error {
	return nil
}

// CreateBanner adds banner to database.
func (s *Storage) CreateBanner(ctx context.Context, bannerID string, description string) (string, error) {
	_, err := s.conn.Exec(ctx, "INSERT INTO banners (banner_id,description) VALUES ($1,$2)", bannerID, description)
	if err != nil {
		return "", fmt.Errorf("create banner: %w", err)
	}
	return bannerID, nil
}

// CreateSlot adds banner to database.
func (s *Storage) CreateSlot(ctx context.Context, slotID string, description string) (string, error) {
	_, err := s.conn.Exec(ctx, "INSERT INTO slots (slot_id,description) VALUES ($1,$2)", slotID, description)
	if err != nil {
		return "", fmt.Errorf("create slot: %w", err)
	}
	return slotID, nil
}

// CreateGroup adds banner to database.
func (s *Storage) CreateGroup(ctx context.Context, groupID string, description string) (string, error) {
	_, err := s.conn.Exec(ctx, "INSERT INTO groups (group_id,description) VALUES ($1,$2)", groupID, description)
	if err != nil {
		return "", fmt.Errorf("create group: %w", err)
	}
	return groupID, nil
}

// Addbanner adds banner to database.
func (s *Storage) AddBanner(ctx context.Context, bannerID string, slotID string) error {
	_, err := s.conn.Exec(ctx, "INSERT INTO rotation (banner_id,slot_id) VALUES ($1,$2)", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("add banner to slot: %w", err)
	}
	return nil
}

// DeleteBanner deleted banner from database.
func (s *Storage) DeleteBanner(ctx context.Context, bannerID string, slotID string) error {
	res, err := s.conn.Exec(ctx, "DELETE FROM rotation WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("delete banner from slot: %w", err)
	}
	rows, err := res.RowsAffected(), err
	if err != nil {
		return fmt.Errorf("get numder of affected rows: %w", err)
	}

	if rows <= 0 {
		return fmt.Errorf("rows are not affected: %w", err)
	}

	return nil
}

// AddClick adds click to database.
func (s *Storage) AddClick(ctx context.Context, bannerID string, slotID string, groupID string, time time.Time) error {
	_, err := s.conn.Exec(ctx, "INSERT INTO clicks (banner_id, slot_id, group_id, time) VALUES ($1,$2,$3, $4)", bannerID, slotID, groupID, time)
	if err != nil {
		return fmt.Errorf("add banner click: %w", err)
	}

	return nil
}

// AddBannerDisplay adds banner display to database.
func (s *Storage) AddBannerDisplay(ctx context.Context, bannerID, slotID string, groupID string, time time.Time) error {
	_, err := s.conn.Exec(ctx, "INSERT INTO displays (banner_id, slot_id, group_id, time) VALUES ($1,$2,$3,$4)", bannerID, slotID, groupID, time)
	if err != nil {
		return fmt.Errorf("add banner display: %w", err)
	}

	return nil
}

// GetBannerDisplay adds banner display to database.
func (s *Storage) GetBannerDisplay(ctx context.Context, slotID string) (bannerDisplay []models.BannerDisplay, err error) {
	_, err = s.conn.Query(ctx, "SELECT * FROM displays WHERE slot_id=$1", &bannerDisplay, slotID)
	if err != nil {
		return nil, fmt.Errorf("get displayed banners: %w", err)
	}

	return bannerDisplay, nil
}

// GetBannerNotDisplay adds banner display to database.
func (s *Storage) GetBannerNotDisplay(ctx context.Context, slotID string) (bannerNotDisplay []models.BannerNotDisplay, err error) {
	_, err = s.conn.Query(ctx,
		"SELECT slot_id, banner_id FROM rotation WHERE slot_id=$1 EXCEPT SELECT slot_id, banner_id FROM displays",
		&bannerNotDisplay, slotID)
	if err != nil {
		return nil, fmt.Errorf("get not displayed banners: %w", err)
	}

	return bannerNotDisplay, nil
}

// GetBannerClick adds banner display to database.
func (s *Storage) GetBannerClick(ctx context.Context, slotID string) (newClicks []models.NewClick, err error) {
	_, err = s.conn.Query(ctx, "SELECT * FROM clicks WHERE slot_id=$1", &newClicks, slotID)
	if err != nil {
		return nil, fmt.Errorf("get banners clicks: %w", err)
	}

	return newClicks, nil
}

// GetBannerInSlot adds banner display to database.
func (s *Storage) GetBannerInSlot(ctx context.Context, slotID string) (rotation []models.Rotation, err error) {
	_, err = s.conn.Query(ctx, "SELECT * FROM rotation WHERE slot_id=$1", &rotation, slotID)
	if err != nil {
		return nil, fmt.Errorf("get banners from slot: %w", err)
	}

	return rotation, nil
}
