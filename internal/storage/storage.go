package storage

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

// Connect creates connection to database.
func (s *Storage) Connect(ctx context.Context, connStr string) error {
	dbase, err := sqlx.ConnectContext(ctx, "pgx", connStr)
	if err != nil {
		return fmt.Errorf("Connect to database: %w", err)
	}
	s.db = dbase
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) AddBanner(ctx context.Context, bannerID string, slotID string) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO banner_location (banner_id,slot_id) VALUES ($1,$2)", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("add banner to slot: %w", err)
	}

	return nil
}

func (s *Storage) DeleteBanner(ctx context.Context, bannerID string, slotID string) error {
	result, err := s.db.ExecContext(ctx, "DELETE FROM banners_location WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("delete banner from slot: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get numder of affected rows: %w", err)
	}

	if rows <= 0 {
		return fmt.Errorf("rows are not affected: %w", err)
	}

	return nil
}

func (s *Storage) AddClick(ctx context.Context, bannerID string, slotID string, groupID string) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO clicks (banner_id,slot_id,group_id, time) VALUES ($1,$2,$3)", bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("add banner click: %w", err)
	}

	return nil
}

func (s *Storage) AddBannerDisplay(ctx context.Context, slotID string, groupID string) (bannerID int, err error) {
	_, error := s.db.ExecContext(ctx, "INSERT INTO displays (banner_id, slot_id,group_id, time) VALUES ($1,$2,$3)", slotID, bannerID, groupID)
	if error != nil {
		return 0, fmt.Errorf("add banner display: %w", err)
	}

	return bannerID, nil
}
