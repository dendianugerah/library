package main

import (
	"log"

	"github.com/dendianugerah/library/db"
	"github.com/dendianugerah/library/model"
	"github.com/dendianugerah/library/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/dendianugerah/library/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library API
// @version 1.0
// @description REST API for library management system
// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
    // Initialize DB
    database, err := db.NewPostgresDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto migrate models
    database.AutoMigrate(&model.Book{}, &model.Borrower{}, &model.Loan{})

    // Initialize handlers
    bookHandler := handler.NewBookHandler(database)
    loanHandler := handler.NewLoanHandler(database)
    borrowerHandler := handler.NewBorrowerHandler(database)

    // Setup router
    r := gin.Default()

    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Book routes
    r.POST("/books", bookHandler.CreateBook)
    r.GET("/books", bookHandler.GetBooks)

    // Borrower routes
    r.POST("/borrowers", borrowerHandler.CreateBorrower)
    r.GET("/borrowers", borrowerHandler.GetBorrowers)

    // Loan routes
    r.POST("/loans", loanHandler.CreateLoan)
    r.GET("/loans", loanHandler.GetLoans)
    r.PUT("/loans/:id/return", loanHandler.ReturnBook)

    r.Run(":8080")
}