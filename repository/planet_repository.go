package repository

import (
	"github.com/jmoiron/sqlx"
	"go-swapi-api/model"
)

type PlanetRepository struct {
	db *sqlx.DB
}

func NewPlanetRepository(db *sqlx.DB) *PlanetRepository {
	return &PlanetRepository{db: db}
}

func (r *PlanetRepository) SavePlanet(planet *model.Planet) error {
	_, err := r.db.Exec(`INSERT INTO planeta (name, diameter, rotation_period) VALUES ($1, $2, $3)`,
		planet.Name, planet.Diameter, planet.RotationPeriod)
	return err
}
