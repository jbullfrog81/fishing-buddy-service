package stores

import (
	"context"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores/enums"
	"github.com/jmoiron/sqlx"
)

const (
	TABLE_IFISH_USERS   = "users"
	TABLE_IFISH_SPECIES = "fish_species"
	TABLE_IFISH_CATCHES = "catches"
)

type IfishStore struct {
	Db *sqlx.DB
}

func NewIfishStore(db *sqlx.DB) *IfishStore {
	return &IfishStore{
		Db: db,
	}
}

type Coordinates struct {
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
}

func (s *IfishStore) newCatch(ctx context.Context, fishSpeciesId enums.FishSpeciesId, fishermanId int, coordinates Coordinates) error {

	query := "INSERT INTO " + TABLE_IFISH_CATCHES + " (user_id, fish_species_id, coordinates, caught_at) VALUES (?, ?, POINT(?, ?), NOW())"

	if err := ctx.Err(); err == context.Canceled {
		return err
	}

	_, err := s.Db.Exec(query, fishermanId, fishSpeciesId, coordinates.Longitude, coordinates.Latitude)
	if err != nil {
		return err
	}

	return nil

}

func (s *IfishStore) NewCatch(ctx context.Context, fishSpeciesId enums.FishSpeciesId, fishermanId int, coordinates Coordinates) error {
	return s.newCatch(ctx, fishSpeciesId, fishermanId, coordinates)
}
