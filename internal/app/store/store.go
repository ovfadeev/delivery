package store

import (
	"database/sql"
	"delivery/internal/app/model/repository"
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
