package product

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Product struct {
	Id          string
	Name        string
	Description string
	CategoryId  string
	CreatedAt   time.Time
}

func (product *Product) BeforeCreate(*gorm.DB) error {
	product.Id = uuid.New().String()
	return nil
}