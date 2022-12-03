package service

import (
	"errors"
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type TransactionService interface {
	CreateTransaction(id_user int, dataInput input.TransactionHistoryCreateInput) (entity.TransactionHistory, error)
	GetUserTransactions(id_user int) ([]entity.TransactionHistory, error)
	GetAllTransactions() ([]entity.TransactionHistory, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
	productRepository     repository.ProductRepository
	categoryRepository    repository.CategoryRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, userRepository repository.UserRepository, productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository) *transactionService {
	return &transactionService{transactionRepository, userRepository, productRepository, categoryRepository}
}

func (s *transactionService) CreateTransaction(id_user int, input input.TransactionHistoryCreateInput) (entity.TransactionHistory, error) {
	productData, err := s.productRepository.FindById(input.ProductID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}
	if productData.ID == 0 {
		return entity.TransactionHistory{}, errors.New("product not found")
	}
	if productData.Stock < input.Quantity {
		return entity.TransactionHistory{}, errors.New("stock not enough")
	}

	userData, err := s.userRepository.FindById(id_user)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	totalPrice := productData.Price * input.Quantity
	if userData.Balance < totalPrice {
		return entity.TransactionHistory{}, errors.New("balance not enough")
	}

	productData.Stock = productData.Stock - input.Quantity
	_, err = s.productRepository.Update(productData.ID, productData)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	userData.Balance = userData.Balance - totalPrice
	_, err = s.userRepository.Update(userData.ID, userData)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	categoryData, err := s.categoryRepository.FindById(productData.CategoryID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}
	categoryData.SoldProductAmount += input.Quantity
	_, err = s.categoryRepository.Update(categoryData.ID, categoryData)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	transaction := entity.TransactionHistory{
		ProductID:  productData.ID,
		UserID:     userData.ID,
		Quantity:   input.Quantity,
		TotalPrice: totalPrice,
	}

	newTransaction, err := s.transactionRepository.Save(transaction)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	return s.transactionRepository.FindById(newTransaction.ID)
}

func (s *transactionService) GetUserTransactions(id_user int) ([]entity.TransactionHistory, error) {
	return s.transactionRepository.FindByUserID(id_user)
}

func (s *transactionService) GetAllTransactions() ([]entity.TransactionHistory, error) {
	return []entity.TransactionHistory{}, nil
}
