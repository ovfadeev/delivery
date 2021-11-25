package repository

import (
	"database/sql"
	"delivery/internal/app/model"
	"reflect"
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



func (r *PointRepository) getByOneParam(param string, value string) (map[int]struct{ model.Point }, error) {
	l := make(map[int]struct{ model.Point })
	rows, err := r.DB.Query("SELECT * FROM "+tablePoint+" WHERE "+param+" = $1", value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		p := model.Point{}
		s := reflect.ValueOf(&p).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for n := 0; i < numCols; n++ {
			field := s.Field(i)
			columns[n] = field.Addr().Interface()
		}
		rows.Scan(columns...)
		l[i] = struct{ model.Point }{p}
	}

	return l, nil
}

func (r *PointRepository) GetByCity(city string) (map[int]struct{ model.Point }, error) {
	return r.getByOneParam("city", city)
}

func (r *PointRepository) GetByZip(zip string) (map[int]struct{ model.Point }, error) {
	return r.getByOneParam("zip", zip)
}
