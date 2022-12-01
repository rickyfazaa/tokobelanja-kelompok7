package service

import (
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
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
