package handler

import (
	"github.com/dendianugerah/library/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
    db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
    return &BookHandler{db: db}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.Book true "Book object"
// @Success 201 {object} model.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
    var book model.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if result := h.db.Create(&book); result.Error != nil {
        c.JSON(400, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(201, book)
}

// GetBooks godoc
// @Summary List all books
// @Description Get all books in the library
// @Tags books
// @Produce json
// @Success 200 {array} model.Book
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
    var books []model.Book
    h.db.Find(&books)
    c.JSON(200, books)
}
