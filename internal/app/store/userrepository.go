package store

import (
	"delivery/internal/app/model"
)

type UserRepository struct {
	store *Store
}

var table = "users"

func (r *UserRepository) Create(u *model.Users) (*model.Users, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO "+table+" (login) VALUES ($1) RETURNING id", u.Login,
	).Scan(u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

// Update create new apikey
func (r *UserRepository) Update(u *model.Users) (*model.Users, error) {
	if err := r.store.db.QueryRow(
		"UPDATE "+table+" SET login = $1 WHERE login = $1", u.Login,
	).Scan(u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetByLogin(login string) (*model.Users, error) {
	if err := r.store.db.QueryRow(
		"SELECT id, created_at, updated_at, login, apikey FROM "+table+"WHERE login = $1", login,
	).Scan(); err != nil {
		return nil, err
	}

	return nil, nil
}
