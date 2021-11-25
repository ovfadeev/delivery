package repository

import (
	"database/sql"
	"delivery/internal/app/model"
)

type PointRepository struct {
	DB *sql.DB
}

var tablePoint = "points"

func (r *PointRepository) Create(p *model.Point) error {
	return r.DB.QueryRow(
		"INSERT INTO "+tablePoint+" (...) VALUES (...) RETURNING id", "",
	).Scan(
		&p.Id,
	)
}

func (r *PointRepository) Update(p *model.Point) error {
	return r.DB.QueryRow(
		"UPDATE "+tablePoint+" SET ... WHERE ... RETURNING id", "",
	).Scan(
		&p.Id,
	)
}

func (r *PointRepository) GetByCity(city string) map[int]struct{ model.Point } {
	//...
	return nil
}

func (r *PointRepository) GetByZip(zip string) map[int]struct{ model.Point } {
	//...
	return nil
}
