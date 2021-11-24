package repository

import (
	"delivery/internal/app/model"
	"delivery/internal/app/store"
	"fmt"
)

type UserRepository struct {
	Store *store.Store
}

var table = "users"

func (r *UserRepository) Create(u *model.Users) error {
	return r.Store.Db.QueryRow(
		"INSERT INTO "+table+" (login) VALUES ($1) RETURNING id", u.Login,
	).Scan(
		u.Id,
	)
}

// Update create new apikey
func (r *UserRepository) Update(u *model.Users) error {
	return r.Store.Db.QueryRow(
		"UPDATE "+table+" SET login = $1 WHERE login = $1", u.Login,
	).Scan(
		&u.Id,
	)
}

func (r *UserRepository) GetByLogin(login string) (*model.Users, error) {
	u := &model.Users{}
	if err := r.Store.Db.QueryRow(
		"SELECT id, created_at, updated_at, login, apikey FROM "+table+" WHERE login = $1", login,
	).Scan(
		&u.Id,
		&u.Create,
		&u.Update,
		&u.Login,
		&u.Apikey,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetByLoginKey(login string, key string) (*model.Users, error) {
	u := &model.Users{}
	if err := r.Store.Db.QueryRow(
		fmt.Sprintf("SELECT id FROM %s WHERE login = $1, apikey = $2", table),
		login,
		key,
	).Scan(
		&u.Id,
	); err != nil {
		return nil, err
	}

	return u, nil
}
