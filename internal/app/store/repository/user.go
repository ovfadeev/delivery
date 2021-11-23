package repository

import (
	"delivery/internal/app/model"
	"delivery/internal/app/store"
)

type UserRepository struct {
	Store *store.Store
}

var table = "users"

func (r *UserRepository) Create(u *model.Users) (*model.Users, error) {
	if err := r.Store.Db.QueryRow("INSERT INTO "+table+" (login) VALUES ($1) RETURNING id", u.Login).Scan(u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Update(u *model.Users) (*model.Users, error) {
	r.Store.Db.QueryRow("UPDATE "+table+" SET ")
}

func (r *UserRepository) GetByLogin(login string) (*model.Users, error) {

	return nil, nil
}
