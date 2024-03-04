package services

import (
	"github.com/Ferriem/gorm/code/datamodels"
	"github.com/Ferriem/gorm/code/repositories"
)

type IUserService interface {
	InsertUser(*datamodels.User) (uint, error)
	InsertUsers([]*datamodels.User) (int64, error)
	GetUserByID(uint) (*datamodels.User, error)
	GetAll() ([]*datamodels.User, error)
	UpdateUser(*datamodels.User) error
	DeleteUser(uint) error
}

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) IUserService {
	return &UserService{userRepository: repository}
}

func (u UserService) InsertUser(user *datamodels.User) (uint, error) {
	return u.userRepository.CreateUser(user)
}

// SkipHook = true
func (u UserService) InsertUsers(users []*datamodels.User) (int64, error) {
	return u.userRepository.CreateUsers(users)
}

func (u UserService) GetUserByID(id uint) (*datamodels.User, error) {
	return u.userRepository.SelectByID(id)
}

func (u UserService) GetAll() ([]*datamodels.User, error) {
	return u.userRepository.SelectAll()
}

// Simply save()
func (u UserService) UpdateUser(user *datamodels.User) error {
	return u.userRepository.Update(user)
}

func (u UserService) DeleteUser(id uint) error {
	return u.userRepository.Delete(id)
}
