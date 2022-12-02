package service

import (
	"errors"
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type CategoryService interface {
	CreateCategory(role_user string, input input.CategoryCreateInput) (entity.Category, error)
	GetAllCategories(role_user string) ([]entity.Category, error)
	GetProductsByCategoryID(id_category int) ([]entity.Product, error)
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

func (s *categoryService) GetAllCategories(role_user string) ([]entity.Category, error) {
	if role_user != "admin" {
		return []entity.Category{}, errors.New("you are not admin")
	}

	return s.categoryRepository.FindAll()
}

func (s *categoryService) GetProductsByCategoryID(id_category int) ([]entity.Product, error) {
	return s.categoryRepository.FindAllProductsByCategoryID(id_category)
}

func (s *categoryService) PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error) {
	if role_user != "admin" {
		return entity.Category{}, errors.New("you are not admin")
	}

	category := entity.Category{
		Type: input.Type,
	}

	_, err := s.categoryRepository.Update(id_category, category)
	if err != nil {
		return entity.Category{}, err
	}

	categoryData, err := s.categoryRepository.FindById(id_category)
	if err != nil {
		return entity.Category{}, err
	}
	if categoryData.ID == 0 {
		return entity.Category{}, errors.New("category not found")
	}

	return categoryData, nil
}

func (s *categoryService) DeleteCategory(role_user string, id_category int) error {
	if role_user != "admin" {
		return errors.New("you are not admin")
	}

	categoryData, err := s.categoryRepository.FindById(id_category)
	if err != nil {
		return err
	}
	if categoryData.ID == 0 {
		return errors.New("category not found")
	}

	return s.categoryRepository.Delete(id_category)
}
