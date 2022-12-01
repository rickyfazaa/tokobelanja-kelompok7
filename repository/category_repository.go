package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	FindAllProductsByCategoryID(id_category int) ([]entity.Product, error)
	FindById(id_category int) (entity.Category, error)
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
	return categories, err
}

func (r *categoryRepository) FindAllProductsByCategoryID(id_category int) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Where("category_id = ?", id_category).Find(&products).Error
	return products, err
}

func (r *categoryRepository) FindById(id_category int) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", id_category).Find(&category).Error
	return category, err
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
