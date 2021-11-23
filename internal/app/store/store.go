package store

import (
	"database/sql"
	"delivery/internal/app/store/repository"
	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	Db             *sql.DB
	userRepository *repository.UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DbUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db

	return nil
}

func (s *Store) Close() error {
	return nil
}

func (s *Store) User() *repository.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &repository.UserRepository{
		Store: s,
	}

	return s.userRepository
}
