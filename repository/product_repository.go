package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product entity.Product) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindAllByCategoryId(id_category int) ([]entity.Product, error)
	FindById(id_product int) (entity.Product, error)
	Update(id_product int, product entity.Product) (entity.Product, error)
	Delete(id_product int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product entity.Product) (entity.Product, error) {
	err := r.db.Preload("Category").Preload("Category").Create(&product).Error
	return product, err
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) FindAllByCategoryId(id_category int) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Preload("Category").Where("category_id = ?", id_category).Find(&products).Error
	return products, err
}

func (r *productRepository) FindById(id_product int) (entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Category").Where("id = ?", id_product).Find(&product).Error
	return product, err
}

func (r *productRepository) Update(id_product int, product entity.Product) (entity.Product, error) {
	err := r.db.Preload("Category").Where("id = ?", id_product).Updates(&product).Error
	return product, err
}

func (r *productRepository) Delete(id_product int) error {
	var product entity.Product
	err := r.db.Where("id = ?", id_product).Delete(&product).Error
	return err
}
