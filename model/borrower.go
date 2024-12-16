package model

// Borrower represents a library member
type Borrower struct {
    ID       uint   `json:"id" gorm:"primaryKey" example:"1"`
    IDCardNo string `json:"id_card_no" gorm:"unique;not null" example:"ID12345678"`
    Name     string `json:"name" gorm:"not null" example:"John Doe"`
    Email    string `json:"email" gorm:"unique;not null" example:"john@example.com"`
}
