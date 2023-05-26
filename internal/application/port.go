package application

import "ps-user/internal/domain"

type UserService interface {
	Create(user domain.User) (uint64, error)
	FindAllUser() ([]domain.User, error)
	FindById(id uint64) (domain.User, error)
	Update(user domain.User) error
	DeleteById(id uint64) error
	FindByUserAndPassword(userName string) (domain.User, error)
	FindListUsers(IDs string) ([]domain.User, error)
	DeleteListId(IDs string) error
}

type UserRepository interface {
	Create(user domain.User) (uint64, error)
	FindAllUser() ([]domain.User, error)
	FindById(id uint64) (domain.User, error)
	Update(user domain.User) error
	DeleteById(id uint64) error
	FindByUserAndPassword(userName string) (domain.User, error)
	FindListUsers(IDs string) ([]domain.User, error)
	DeleteListId(IDs string) error
}
