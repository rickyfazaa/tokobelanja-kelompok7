package service

import (
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"

	"golang.org/x/crypto/bcrypt"
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "customer",
		Balance:  0,
	}

	return s.userRepository.Save(user)
}

func (s *userService) RegisterAdmin(input input.UserRegisterInput) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) LoginUser(input input.UserLoginInput) (string, error) {
	return "", nil
}

func (s *userService) TopUpUser(input input.UserPatchTopUpInput) error {
	return nil
}
