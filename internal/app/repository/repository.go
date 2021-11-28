package repository

type repository interface {
	Create(i interface{}) error
	Update(i interface{}) error
}
