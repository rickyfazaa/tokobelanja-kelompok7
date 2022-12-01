package service

import "project4/repository"

type TransactionService interface {
	CreateTransaction(input input.TransactionCreateInput) (entity.Transaction, error)
	GetUserTransactions(id_user int) ([]entity.Transaction, error)
	GetAllTransactions() ([]entity.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) CreateTransaction(input input.TransactionCreateInput) (entity.Transaction, error) {
	// TODO
}

func (s *transactionService) GetUserTransactions(id_user int) ([]entity.Transaction, error) {
	// TODO
}

func (s *transactionService) GetAllTransactions() ([]entity.Transaction, error) {
	// TODO
}
