package handler

import (
    "github.com/dendianugerah/library/model"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type BorrowerHandler struct {
    db *gorm.DB
}

func NewBorrowerHandler(db *gorm.DB) *BorrowerHandler {
    return &BorrowerHandler{db: db}
}

// CreateBorrower godoc
// @Summary Create a new borrower
// @Description Register a new library member
// @Tags borrowers
// @Accept json
// @Produce json
// @Param borrower body model.Borrower true "Borrower object"
// @Success 201 {object} model.Borrower
// @Failure 400 {object} map[string]string
// @Router /borrowers [post]
func (h *BorrowerHandler) CreateBorrower(c *gin.Context) {
    var borrower model.Borrower
    if err := c.ShouldBindJSON(&borrower); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if result := h.db.Create(&borrower); result.Error != nil {
        c.JSON(400, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(201, borrower)
}

// GetBorrowers godoc
// @Summary List all borrowers
// @Description Get all registered library members
// @Tags borrowers
// @Produce json
// @Success 200 {array} model.Borrower
// @Router /borrowers [get]
func (h *BorrowerHandler) GetBorrowers(c *gin.Context) {
    var borrowers []model.Borrower
    h.db.Find(&borrowers)
    c.JSON(200, borrowers)
}
