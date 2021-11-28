package repository

type repository interface {
	Create(i interface{}) (interface{}, error)
	Update(i interface{}) (interface{}, error)
}
