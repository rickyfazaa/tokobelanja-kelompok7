package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product entity.Product) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	Update(id_product int, product entity.Product) (entity.Product, error)
	Delete(id_product int, product entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product entity.Product) (entity.Product, error) {
	err := r.db.Preload("Category").Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) Update(id_product int, product entity.Product) (entity.Product, error) {
	err := r.db.Preload("Category").Where("id = ?", id_product).Updates(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Delete(id_product int, product entity.Product) error {
	err := r.db.Where("id = ?", id_product).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
