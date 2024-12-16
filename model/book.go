package model

// Book represents a book in the library
type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey" example:"1"`
    Title  string `json:"title" gorm:"not null" example:"The Go Programming Language"`
    ISBN   string `json:"isbn" gorm:"unique;not null" example:"978-0134190440"`
    Stock  int    `json:"stock" gorm:"not null" example:"5"`
}
