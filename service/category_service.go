package service

import (
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type CategoryService interface {
	CreateCategory(input input.CategoryCreateInput) (entity.Category, error)
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

func (s *categoryService) CreateCategory(input input.CategoryCreateInput) (entity.Category, error) {
	// TODO
}

func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	// TODO
}

func (s *categoryService) PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error) {
	// TODO
}

func (s *categoryService) DeleteCategory(role_user string, id_category int) error {
	// TODO
}
