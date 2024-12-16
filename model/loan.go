package model

import "time"

// Loan represents a book borrowing transaction
type Loan struct {
    ID         uint       `json:"id" gorm:"primaryKey" example:"1"`
    BookID     uint       `json:"book_id" gorm:"not null" example:"1"`
    BorrowerID uint       `json:"borrower_id" gorm:"not null" example:"1"`
    BorrowDate time.Time  `json:"borrow_date" gorm:"not null" example:"2023-01-01T00:00:00Z"`
    DueDate    time.Time  `json:"due_date" gorm:"not null" example:"2023-01-15T00:00:00Z"`
    ReturnDate *time.Time `json:"return_date" example:"2023-01-10T00:00:00Z"`
    IsLate     bool       `json:"is_late" example:"false"`
    Book       Book       `json:"book" gorm:"foreignKey:BookID"`
    Borrower   Borrower   `json:"borrower" gorm:"foreignKey:BorrowerID"`
}
