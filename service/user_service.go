package service

import (
	"errors"
	"tokobelanja-kelompok7/middleware"
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input input.UserRegisterInput) (entity.User, error)
	RegisterAdmin(input input.UserRegisterInput) (entity.User, error)
	LoginUser(userInput input.UserLoginInput) (string, error)
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "admin",
		Balance:  0,
	}

	return s.userRepository.Save(user)
}

func (s *userService) LoginUser(userInput input.UserLoginInput) (string, error) {
	userData, err := s.userRepository.FindByEmail(userInput.Email)
	if err != nil {
		return "", err
	}
	if userData.ID == 0 {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userInput.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	return middleware.GenerateToken(userData.ID, userData.Role)
}

func (s *userService) TopUpUser(input input.UserPatchTopUpInput) error {
	return nil
}
