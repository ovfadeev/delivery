package store

import (
	"database/sql"
	"delivery/internal/app/repository"
	"encoding/json"
	_ "github.com/lib/pq"
)

type Store struct {
	DB    *sql.DB
	user  *repository.UserRepository
	point *repository.PointRepository
}

func New(config *Config) *Store {
	return &Store{}
}

func (s *Store) Open(DBUrl string) error {
	db, err := sql.Open("postgres", DBUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Store) Close() error {
	err := s.DB.Close()
	return err
}

func (s *Store) User() *repository.UserRepository {
	if s.user == nil {
		s.user = &repository.UserRepository{s.DB}
	}

	return s.user
}

func (s *Store) Point() *repository.PointRepository {
	if s.point == nil {
		s.point = &repository.PointRepository{s.DB}
	}

	return s.point
}

func (s *Store) GetUserLogin(login string, key string) (int, error) {
	u, err := s.User().GetByLoginKey(login, key)
	if err != nil {
		return 0, err
	}
	return u.Id, err

}

func (s *Store) GetPointsFromCity(city string) ([]byte, error) {
	p, err := s.Point().GetByCity(city)
	if err != nil {
		return nil, err
	}
	return json.Marshal(p)
}

func (s *Store) GetPointsFromZip(zip string) ([]byte, error) {
	p, err := s.Point().GetByZip(zip)
	if err != nil {
		return nil, err
	}
	return json.Marshal(p)
}
