package repository

type Repository interface {
	GetInfo() (map[string]interface{}, error)
}
