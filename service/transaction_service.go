package service

import (
	"tokobelanja-kelompok7/model/entity"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/repository"
)

type TransactionService interface {
	CreateTransaction(dataInput input.TransactionHistoryCreateInput) (entity.TransactionHistory, error)
	GetUserTransactions(id_user int) ([]entity.TransactionHistory, error)
	GetAllTransactions() ([]entity.TransactionHistory, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) CreateTransaction(input input.TransactionHistoryCreateInput) (entity.TransactionHistory, error) {
	return entity.TransactionHistory{}, nil
}

func (s *transactionService) GetUserTransactions(id_user int) ([]entity.TransactionHistory, error) {
	return []entity.TransactionHistory{}, nil
}

func (s *transactionService) GetAllTransactions() ([]entity.TransactionHistory, error) {
	return []entity.TransactionHistory{}, nil
}
