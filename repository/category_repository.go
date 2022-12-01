package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	Update(id_category int, category entity.Category) (entity.Category, error)
	Delete(id_category int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *categoryRepository) Update(id_category int, category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", id_category).Updates(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id_category int) error {
	err := r.db.Where("id = ?", id_category).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}
