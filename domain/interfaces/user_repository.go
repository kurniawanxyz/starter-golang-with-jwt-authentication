package interfaces

import "github.com/kurniawanxzy/backend-olshop/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindById(id string) (*entities.User, error)
}