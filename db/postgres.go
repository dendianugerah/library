package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
    dsn := "postgresql://username:password@localhost:5432/library_db?sslmode=disable"
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
