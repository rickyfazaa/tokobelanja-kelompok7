package service

import (
	"errors"
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type ProductService interface {
	CreateProduct(role_user string, inputBody input.ProductCreateInput) (entity.Product, error)
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

func (s *productService) CreateProduct(role_user string, inputBody input.ProductCreateInput) (entity.Product, error) {
	if role_user != "admin" {
		return entity.Product{}, errors.New("you are not admin")
	}

	categoryData, err := s.productRepository.FindCategoryByCategoryId(inputBody.CategoryID)
	if err != nil {
		return entity.Product{}, err
	}
	if categoryData.ID == 0 {
		return entity.Product{}, errors.New("category not found")
	}

	product := entity.Product{
		Title:      inputBody.Title,
		Price:      inputBody.Price,
		Stock:      inputBody.Stock,
		CategoryID: inputBody.CategoryID,
	}

	return s.productRepository.Save(product)
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	return s.productRepository.FindAll()
}

func (s *productService) UpdateProduct(id_product int, input input.ProductUpdateInput) (entity.Product, error) {
	return entity.Product{}, nil
}

func (s *productService) DeleteProduct(id_product int) error {
	return nil
}
