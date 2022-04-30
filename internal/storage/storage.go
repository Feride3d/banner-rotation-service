package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Connection to database

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, connStr string) error {
	dbase, err := sqlx.ConnectContext(ctx, "pgx", connStr)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v", err)
	}
	s.db = dbase
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

/* // изучить вариант замены на pool.connect для concurrency
dbUrl := "postgres://postgres:mypassword@localhost:5432/postgres"
dbpool, err := pgxpool.Connect(context.Background(), dbUrl)
if err != nil {
	log.Fatalf("Unable to connect to database: %v", err)
}
defer dbpool.Close() // close DB pool
log.Println("Connected") */

func (s *Storage) AddBanner(ctx context.Context, bannerID string, slotID string) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO banner_location (banner_id,slot_id) VALUES ($1,$2)", bannerID, slotID)
	if err != nil {
		return errors.Wrapf(err, "unable to add banner on slot for rotation, %w")
	}

	return nil
}

func (s *Storage) DeleteBanner(ctx context.Context, bannerID string, slotID string) error {
	result, err := s.db.ExecContext(ctx, "DELETE FROM banners_location WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return errors.Wrapf(err, "unable to delete banner from slot, %w")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "unable to get numder of affected rows, %w")
	}

	if rows <= 0 {
		return errors.Wrapf(err, "rows are not affected, %w")
	}

	return nil
}

func (s *Storage) AddClick(ctx context.Context, bannerID string, slotID string, groupID string) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO clicks (banner_id,slot_id,group_id) VALUES ($1,$2,$3)", bannerID, slotID, groupID)
	if err != nil {
		return errors.Wrapf(err, "unable to add banner click, %w")
	}

	return nil
}

func (s *Storage) AddBannerDisplay(ctx context.Context, slotID string, groupID string) (bannerID int, err error) {
	_, error := s.db.ExecContext(ctx, "INSERT INTO views (banner_id, slot_id,group_id) VALUES ($1,$2,$3)", slotID, bannerID, groupID)
	if error != nil {
		return 0, errors.Wrapf(err, "unable to add banner display, %w")
	}

	return bannerID, nil
}
