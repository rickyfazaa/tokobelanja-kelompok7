package controller

import (
	"tokobelanja-kelompok7/service"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *transactionController {
	return &transactionController{transactionService}
}

func (h *transactionController) CreateTransaction(c *gin.Context) {
	// TODO
}

func (h *transactionController) GetUserTransactions(c *gin.Context) {
	// TODO
}

func (h *transactionController) GetAllTransactions(c *gin.Context) {
	// TODO
}
