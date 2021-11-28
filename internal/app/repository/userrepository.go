package repository

import (
	"database/sql"
	"delivery/internal/app/model"
)

type UserRepository struct {
	DB *sql.DB
}

var tableUser = "users"

func (r *UserRepository) Create(u *model.Users) error {
	return r.DB.QueryRow(
		"INSERT INTO "+tableUser+" (login) VALUES ($1) RETURNING id", u.Login,
	).Scan(
		&u.Id,
	)
}

// Update create new apikey
func (r *UserRepository) Update(u *model.Users) error {
	return r.DB.QueryRow(
		"UPDATE "+tableUser+" SET login = $1 WHERE login = $1 RETURNING id", u.Login,
	).Scan(
		&u.Id,
	)
}

func (r *UserRepository) GetByLogin(login string) (*model.Users, error) {
	u := &model.Users{}
	if err := r.DB.QueryRow(
		"SELECT id, created_at, updated_at, login, apikey FROM "+tableUser+" WHERE login = $1", login,
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
	if err := r.DB.QueryRow(
		"SELECT id FROM "+tableUser+" WHERE login = $1 AND apikey = $2",
		login,
		key,
	).Scan(
		&u.Id,
	); err != nil {
		return nil, err
	}

	return u, nil
}
