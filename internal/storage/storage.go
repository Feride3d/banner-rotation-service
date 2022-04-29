package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jackc/pgx/v4/pgxpool" // The sql package must be used in conjunction with a database driver.
	"github.com/Feride3d/banner-rotation-service/internal/service/mab"
	"github.com/Feride3d/banner-rotation-service/internal/storage/models"
)

	// connection to database 
	func initHandlers(pool *pgxpool.Pool) http.Handler { // the connection pool for concurrency
		r := mux.NewRouter()
	dbpool, err := pgxpool.Connect(context.Background(), dbURL) 
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()
	log.Infof("Connected")


func (s *Storage) AddBanner(bannerID string, slotID string) error {
	_, err := s.dbpool.Exec(context.Background(), "INSERT INTO banner_location (banner_id,slot_id) VALUES ($1,$2)", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("unable to add banner on slot for rotation, %w", err)
	}

	return nil
}

func (s *Storage) DeleteBanner(bannerID string, slotID string) error { 
	result, err := s.dbpool.Exec(context.Background(), "DELETE FROM banners_location WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("unable to delete banner from slot, %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to get numder of affected rows, %w", err)
	}

	if rows <= 0 {
		return fmt.Errorf("rows are not affected, %w", err)
	}

	return nil
}

func (s *Storage) AddClick(bannerID string, slotID string, groupID string) error {
	_, err := s.dbpool.Exec(context.Background(), "INSERT INTO clicks (banner_id,slot_id,group_id) VALUES ($1,$2,$3)", bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("unable to add banner click, %w", err)
	}

	return nil
}

func (s *Storage) AddBannerDisplay(slotID string, groupID string) (mab.Ucb int, err error) {
	_, err := s.dbpool.Exec(context.Background(), "INSERT INTO views (banner_id, slot_id,group_id) VALUES ($1,$2,$3)", slotID, bannerID, groupID)
	if err != nil {
		return fmt.Errorf("unable to add banner display, %w", err)
	}

	return mab.Ucb(), nil
}