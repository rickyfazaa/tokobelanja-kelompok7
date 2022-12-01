package response

import "time"

type TransactionHistoryCreateResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionHistoryGetUserResponse struct {
	ID         int                `json:"id"`
	ProductID  int                `json:"product_id"`
	UserID     int                `json:"user_id"`
	Quantity   int                `json:"quantity"`
	TotalPrice int                `json:"total_price"`
	Product    TransactionProduct `json:"Product"`
}

type TransactionProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TransactionHistoryGetAdminResponse struct {
	ID         int                `json:"id"`
	ProductID  int                `json:"product_id"`
	UserID     int                `json:"user_id"`
	Quantity   int                `json:"quantity"`
	TotalPrice int                `json:"total_price"`
	Product    TransactionProduct `json:"Product"`
	User       TransactionUser    `json:"User"`
}

type TransactionUser struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
