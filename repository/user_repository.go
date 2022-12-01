package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Update(id_user int, user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *userRepository) Update(id_user int, user entity.User) (entity.User, error) {
	err := r.db.Where("id = ?", id_user).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
