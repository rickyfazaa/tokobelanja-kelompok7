package main

import (
	"os"
	"project4/config"
	"project4/controller"
	"project4/middleware"
	"project4/repository"
	"project4/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	db := config.InitDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	router := gin.Default()

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PATCH("/users/topup", middleware.AuthMiddleware, userController.PatchTopUpUser)
	// Create Admin
	router.POST("/users/admin", userController.RegisterAdmin)

	// Categories
	router.POST("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware, categoryController.GetAllCategories)
	router.PATCH("/categories/:id", middleware.AuthMiddleware, categoryController.PatchCategory)
	router.DELETE("/categories/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Products
	router.POST("/products", middleware.AuthMiddleware, productController.CreateProduct)
	router.GET("/products", middleware.AuthMiddleware, productController.GetAllProducts)
	router.PUT("/products/:id", middleware.AuthMiddleware, productController.UpdateProduct)
	router.DELETE("/products/:id", middleware.AuthMiddleware, productController.DeleteProduct)

	// TransactionHistories
	router.POST("/transactions", middleware.AuthMiddleware, transactionController.CreateTransaction)
	router.GET("/transactions/my-transactions", middleware.AuthMiddleware, transactionController.GetUserTransactions)
	router.GET("/transactions/user-transactions", middleware.AuthMiddleware, transactionController.GetAllTransactions)

	router.Run(":" + port)
}
