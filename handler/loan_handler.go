package handler

import (
	"time"

	"github.com/dendianugerah/library/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoanHandler struct {
    db *gorm.DB
}

func NewLoanHandler(db *gorm.DB) *LoanHandler {
    return &LoanHandler{db: db}
}

// CreateLoan godoc
// @Summary Create a new loan
// @Description Create a new book loan
// @Tags loans
// @Accept json
// @Produce json
// @Param loan body model.Loan true "Loan object"
// @Success 201 {object} model.Loan
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /loans [post]
func (h *LoanHandler) CreateLoan(c *gin.Context) {
    var loan model.Loan
    if err := c.ShouldBindJSON(&loan); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Check if borrower has active loans
    var activeLoan model.Loan
    if result := h.db.Where("borrower_id = ? AND return_date IS NULL", loan.BorrowerID).First(&activeLoan); result.Error == nil {
        c.JSON(400, gin.H{"error": "Borrower already has an active loan"})
        return
    }

    // Check book availability
    var book model.Book
    if result := h.db.First(&book, loan.BookID); result.Error != nil {
        c.JSON(400, gin.H{"error": "Book not found"})
        return
    }

    if book.Stock <= 0 {
        c.JSON(400, gin.H{"error": "Book out of stock"})
        return
    }

    // Validate loan duration
    loan.BorrowDate = time.Now()
    if loan.DueDate.Sub(loan.BorrowDate) > 30*24*time.Hour {
        c.JSON(400, gin.H{"error": "Loan duration cannot exceed 30 days"})
        return
    }

    // Create loan and update stock
    tx := h.db.Begin()
    if err := tx.Create(&loan).Error; err != nil {
        tx.Rollback()
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := tx.Model(&book).Update("stock", book.Stock-1).Error; err != nil {
        tx.Rollback()
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    tx.Commit()
    c.JSON(201, loan)
}

// ReturnBook godoc
// @Summary Return a book
// @Description Mark a book as returned
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Success 200 {object} model.Loan
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /loans/{id}/return [put]
func (h *LoanHandler) ReturnBook(c *gin.Context) {
    loanID := c.Param("id")
    var loan model.Loan
    
    if result := h.db.First(&loan, loanID); result.Error != nil {
        c.JSON(404, gin.H{"error": "Loan not found"})
        return
    }

    if loan.ReturnDate != nil {
        c.JSON(400, gin.H{"error": "Book already returned"})
        return
    }

    now := time.Now()
    loan.ReturnDate = &now
    loan.IsLate = now.After(loan.DueDate)

    tx := h.db.Begin()
    if err := tx.Save(&loan).Error; err != nil {
        tx.Rollback()
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := tx.Model(&loan.Book).Update("stock", loan.Book.Stock+1).Error; err != nil {
        tx.Rollback()
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    tx.Commit()
    c.JSON(200, loan)
}

// GetLoans godoc
// @Summary Get all loans
// @Description Get all loans with their status (on-time/late)
// @Tags loans
// @Produce json
// @Param status query string false "Filter by status (late/ontime)"
// @Success 200 {array} model.Loan
// @Router /loans [get]
func (h *LoanHandler) GetLoans(c *gin.Context) {
    var loans []model.Loan
    query := h.db.Preload("Book").Preload("Borrower")
    
    // Filter by status if provided
    status := c.Query("status")
    if status == "late" {
        query = query.Where("is_late = ?", true)
    } else if status == "ontime" {
        query = query.Where("is_late = ?", false)
    }

    query.Find(&loans)
    c.JSON(200, loans)
}
