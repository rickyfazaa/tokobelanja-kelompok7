package service

import (
	"errors"
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type CategoryService interface {
	CreateCategory(role_user string, input input.CategoryCreateInput) (entity.Category, error)
	GetAllCategories() ([]entity.Category, error)
	PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error)
	DeleteCategory(role_user string, id_category int) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(role_user string, input input.CategoryCreateInput) (entity.Category, error) {
	if role_user != "admin" {
		return entity.Category{}, errors.New("you are not admin")
	}

	category := entity.Category{
		Type:              input.Type,
		SoldProductAmount: 0,
	}

	return s.categoryRepository.Save(category)
}

func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	return []entity.Category{}, nil
}

func (s *categoryService) PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error) {
	return entity.Category{}, nil
}

func (s *categoryService) DeleteCategory(role_user string, id_category int) error {
	return nil
}
