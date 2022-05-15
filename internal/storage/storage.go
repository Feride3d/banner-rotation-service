package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	conn *pgxpool.Pool
}

func NewStorage(connStr string) (*Storage, error) {
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

// Addbanner adds banner to database.
func (s *Storage) AddBanner(ctx context.Context, bannerID string, slotID string) error {
	_, err := s.conn.Exec(ctx, "INSERT INTO banner_location (banner_id,slot_id) VALUES ($1,$2)", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("add banner to slot: %w", err)
	}
	return nil
}

// DeleteBanner deleted banner from database.
func (s *Storage) DeleteBanner(ctx context.Context, bannerID string, slotID string) error {
	result, err := s.conn.Exec(ctx, "DELETE FROM banners_location WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("delete banner from slot: %w", err)
	}
	rows, err := result.RowsAffected(), err
	if err != nil {
		return fmt.Errorf("get numder of affected rows: %w", err)
	}

	if rows <= 0 {
		return fmt.Errorf("rows are not affected: %w", err)
	}

	return nil
}

// AddClick adds click to database.
func (s *Storage) AddClick(ctx context.Context, bannerID string, slotID string, groupID string) error {
	_, err := s.conn.Exec(ctx, "INSERT INTO clicks (banner_id,slot_id,group_id, time) VALUES ($1,$2,$3)", bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("add banner click: %w", err)
	}

	return nil
}

// AddBannerDisplay adds banner display to database.
func (s *Storage) AddBannerDisplay(ctx context.Context, slotID string, groupID string) (bannerID int, err error) {
	_, error := s.conn.Exec(ctx, "INSERT INTO displays (banner_id, slot_id,group_id, time) VALUES ($1,$2,$3)", slotID, bannerID, groupID)
	if error != nil {
		return 0, fmt.Errorf("add banner display: %w", err)
	}

	return bannerID, nil
}
