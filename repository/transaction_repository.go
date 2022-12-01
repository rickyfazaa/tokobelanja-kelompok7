package repository

import (
	"tokobelanja-kelompok7/model/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction entity.TransactionHistory) (entity.TransactionHistory, error)
	FindByUserID(id_user int) ([]entity.TransactionHistory, error)
	FindAll() ([]entity.TransactionHistory, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Save(transaction entity.TransactionHistory) (entity.TransactionHistory, error) {
	err := r.db.Preload("Product").Preload("User").Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) FindByUserID(id_user int) ([]entity.TransactionHistory, error) {
	var transactions []entity.TransactionHistory
	err := r.db.Preload("Product").Preload("User").Where("id_user = ?", id_user).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *transactionRepository) FindAll() ([]entity.TransactionHistory, error) {
	var transactions []entity.TransactionHistory
	err := r.db.Preload("Product").Preload("User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
