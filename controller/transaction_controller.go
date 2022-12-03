package controller

import (
	"net/http"
	"tokobelanja-kelompok7/helper"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/model/response"
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
	var inputBody input.TransactionHistoryCreateInput

	id_user := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	transactionData, err := h.transactionService.CreateTransaction(id_user, inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	transactionResponse := response.TransactionHistoryCreateResponse{
		Message: "Transaction Success",
		TransactionBill: response.TransactionBill{
			TotalPrice:   transactionData.TotalPrice,
			Quantity:     transactionData.Quantity,
			ProductTitle: transactionData.Product.Title,
		},
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			transactionResponse,
		),
	)
}

func (h *transactionController) GetUserTransactions(c *gin.Context) {
	var allTransactions []response.TransactionHistoryGetUserResponse

	id_user := c.MustGet("currentUser").(int)

	transactionData, err := h.transactionService.GetUserTransactions(id_user)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	for _, transaction := range transactionData {
		allTransactionsTmp := response.TransactionHistoryGetUserResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: response.TransactionProduct{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      transaction.Product.Stock,
				CategoryID: transaction.Product.CategoryID,
				CreatedAt:  transaction.Product.CreatedAt,
				UpdatedAt:  transaction.Product.UpdatedAt,
			},
		}
		allTransactions = append(allTransactions, allTransactionsTmp)
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			allTransactions,
		),
	)
}

func (h *transactionController) GetAllTransactions(c *gin.Context) {
	// TODO
}
