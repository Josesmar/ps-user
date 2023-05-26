package service

import (
	"errors"
	"ps-user/internal/application"
	"ps-user/internal/domain"
)

type userService struct {
	userRepository application.UserRepository
}

func NewUserService(userRepository application.UserRepository) application.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Create implements application.UserService.
func (u *userService) Create(user domain.User) (uint64, error) {
	userId, err := u.userRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// DeleteById implements application.UserService.
func (u *userService) DeleteById(id uint64) error {
	err := u.userRepository.DeleteById(id)
	if err != nil {
		return errors.New("error deleting user")
	}
	return nil
}

// DeleteListId implements application.UserService.
func (u *userService) DeleteListId(IDs string) error {
	err := u.userRepository.DeleteListId(IDs)
	if err != nil {
		return errors.New("error deleting user")
	}
	return nil
}

// FindAllUser implements application.UserService.
func (u *userService) FindAllUser() ([]domain.User, error) {
	users, err := u.userRepository.FindAllUser()
	if err != nil {
		return users, errors.New("error finding user")
	}
	return users, nil
}

// FindById implements application.UserService.
func (u *userService) FindById(id uint64) (domain.User, error) {
	users, err := u.userRepository.FindById(id)
	if err != nil {
		return users, errors.New("error finding user")
	}
	return users, nil
}

// FindByUserAndPassword implements application.UserService.
func (u *userService) FindByUserAndPassword(userName string) (domain.User, error) {
	users, err := u.userRepository.FindByUserAndPassword(userName)
	if err != nil {
		return users, errors.New("error finding user")
	}
	return users, nil
}

// FindListUsers implements application.UserService.
func (u *userService) FindListUsers(IDs string) ([]domain.User, error) {
	users, err := u.userRepository.FindListUsers(IDs)
	if err != nil {
		return users, errors.New("error finding user")
	}
	return users, nil

}

// Update implements application.UserService.
func (u *userService) Update(user domain.User) error {
	err := u.userRepository.Update(user)
	if err != nil {
		return errors.New("error updating user")
	}
	return nil
}
