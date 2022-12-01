package service

import (
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type ProductService interface {
	CreateProduct(input input.ProductCreateInput) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(id_product int, input input.ProductUpdateInput) (entity.Product, error)
	DeleteProduct(id_product int) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(input input.ProductCreateInput) (entity.Product, error) {
	// TODO
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	// TODO
}

func (s *productService) UpdateProduct(id_product int, input input.ProductUpdateInput) (entity.Product, error) {
	// TODO
}

func (s *productService) DeleteProduct(id_product int) error {
	// TODO
}
