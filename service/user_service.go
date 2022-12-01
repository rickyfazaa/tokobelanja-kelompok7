package service

import (
	"project4/model/entity"
	"project4/model/input"
	"project4/repository"
)

type UserService interface {
	RegisterUser(input input.UserRegisterInput) (entity.User, error)
	RegisterAdmin(input input.UserRegisterInput) (entity.User, error)
	LoginUser(input input.UserLoginInput) (string, error)
	TopUpUser(input input.UserPatchTopUpInput) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(input input.UserRegisterInput) (entity.User, error) {
	// TODO
}

func (s *userService) RegisterAdmin(input input.UserRegisterInput) (entity.User, error) {
	// TODO
}

func (s *userService) LoginUser(input input.UserLoginInput) (string, error) {
	// TODO
}

func (s *userService) TopUpUser(input input.UserPatchTopUpInput) error {
	// TODO
}
