package productCategory

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	Id   string
	Name string
}

func (productCategory *ProductCategory) BeforeCreate(*gorm.DB) error {
	productCategory.Id = uuid.New().String()
	return nil
}